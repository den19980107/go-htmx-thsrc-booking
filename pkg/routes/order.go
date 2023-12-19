package routes

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/ordervalidation"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/models"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/station"
)

type (
	orderController struct {
		controller.Controller
	}

	orderPageData struct {
		StationList []station.StationCode
	}

	orderForm struct {
		StartTime        string `form:"start-time" validate:"required"`
		EndTime          string `form:"end-time" validate:"required"`
		DepartureStation string `form:"departure-station" validate:"required"`
		ArrivalStation   string `form:"arrival-station" validate:"required"`
		IdNumber         string `form:"id-number" validate:"required"`
		PhoneNumber      string `form:"phone-number" validate:"required"`
		Email            string `form:"email" validate:"required,email"`
		Submission       controller.FormSubmission
	}

	orderValidationData struct {
		ImageUrl string
	}

	orderValidationForm struct {
		OrderId    int    `form:"order-id" validate:"required"`
		JsessionId string `form:"jsession-id" validate:"required"`
		Captcha    string `form:"captcha" validate:"required"`
		Cookies    string `form:"cookies" validate:"required"`
		Submission controller.FormSubmission
	}
)

func (c *orderController) Get(ctx echo.Context) error {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "order"
	page.Title = "Order"
	page.Data = orderPageData{
		StationList: station.GetAllStations(),
	}
	form := orderForm{
		StartTime: now.Format("2006-01-02T15:04"),
		EndTime:   now.Add(time.Hour).Format("2006-01-02T15:04"),
	}
	user, err := c.Container.Auth.GetAuthenticatedUser(ctx)
	if err == nil {
		form.Email = user.Email
	}
	page.Form = form

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*orderForm)
		log.Printf("form: %+v", page.Form)
	}

	return c.RenderPage(ctx, page)
}

func (c *orderController) GetValidate(ctx echo.Context) error {
	validateId, err := strconv.Atoi(ctx.Param("validateId"))
	if err != nil {
		msg.Warning(ctx, "validate id format not correct!")
		return c.Redirect(ctx, "home")
	}

	currentUser, err := c.Container.Auth.GetAuthenticatedUser(ctx)
	if err != nil {
		msg.Warning(ctx, "you have not login yet!")
		return c.Redirect(ctx, "login")
	}

	orderValidation, err := c.Container.ORM.OrderValidation.Query().Where(ordervalidation.ID(validateId)).Only(ctx.Request().Context())
	if err != nil {
		msg.Warning(ctx, "validate id not found!")
		return c.Redirect(ctx, "home")
	}

	order, err := orderValidation.QueryOrder().Only(ctx.Request().Context())
	if err != nil {
		msg.Warning(ctx, "validate order not found!")
		return c.Redirect(ctx, "home")
	}

	orderUser, err := order.QueryUser().Only(ctx.Request().Context())
	if err != nil {
		msg.Warning(ctx, "order's user not found!")
		return c.Redirect(ctx, "home")
	}

	if orderUser.ID != currentUser.ID {
		msg.Warning(ctx, "order is not blongs to you!")
		return c.Redirect(ctx, "home")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "order-validation"
	page.Title = "Order validation"
	page.Data = orderValidationData{
		ImageUrl: orderValidation.CaptchaImage,
	}
	page.Form = orderValidationForm{
		OrderId:    order.ID,
		JsessionId: orderValidation.JessionID,
		Captcha:    "",
		Cookies:    orderValidation.Cookies,
		Submission: controller.FormSubmission{},
	}
	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*orderForm)
		log.Printf("form: %+v", page.Form)
	}

	return c.RenderPage(ctx, page)
}

func (c *orderController) Post(ctx echo.Context) error {
	var form orderForm
	ctx.Set(context.FormKey, &form)

	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse create order form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	startTime, err := time.ParseInLocation("2006-01-02T15:04", form.StartTime, time.Local)
	if err != nil {
		form.Submission.SetFieldError("StartTime", "unable to parse start time")
		return c.Get(ctx)
	}
	endTime, err := time.ParseInLocation("2006-01-02T15:04", form.EndTime, time.Local)
	if err != nil {
		form.Submission.SetFieldError("EndTime", "unable to parse end time")
		return c.Get(ctx)
	}
	if startTime.Year() != endTime.Year() || startTime.Month() != endTime.Month() || startTime.Day() != endTime.Day() {
		form.Submission.SetFieldError("StartTime", "start time and end time should be at the same day")
		form.Submission.SetFieldError("EndTime", "start time and end time should be at the same day")
		return c.Get(ctx)
	}

	user, err := c.Container.Auth.GetAuthenticatedUser(ctx)
	if err != nil {
		return c.Redirect(ctx, "login")
	}

	o, err := c.Container.ORM.Order.
		Create().
		SetUser(user).
		SetStartTime(startTime).
		SetEndTime(endTime).
		SetEmail(form.Email).
		SetPhoneNumber(form.PhoneNumber).
		SetArrivalStation(form.ArrivalStation).
		SetDepartureStation(form.DepartureStation).
		SetIDNumber(form.IdNumber).
		SetErrorMessage("").
		Save(ctx.Request().Context())

	if err != nil {
		// ctx.Logger().Infof("create order failed, err: %s", err)
		log.Printf("create order failed, err: %s", err)
		return c.Fail(err, "unable to create order")
	}

	log.Printf("order from %s to %s at %s is created!", o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04"))
	msg.Success(ctx, fmt.Sprintf("Order from %s to %s at %s is created!", o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04")))
	return c.Redirect(ctx, "home")
}

func (c *orderController) PostValidation(ctx echo.Context) error {
	var form orderValidationForm
	ctx.Set(context.FormKey, &form)

	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse order validate form")
	}

	o, err := c.Container.ORM.Order.Query().Where(order.ID(form.OrderId)).Only(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "failed to find order by order id")
	}

	cw, err := c.Container.CrawlerStore.Get(o.ID)
	if err != nil {
		_, err = o.Update().SetStatus(models.OrderFailed).SetErrorMessage(err.Error()).Save(ctx.Request().Context())
		return c.Fail(err, "failed to find the crawler")
	}
	defer c.Container.CrawlerStore.Delete(o.ID)

	err = cw.CompleteOrder(*o, form.Captcha, form.JsessionId)
	if err != nil {
		_, err = o.Update().SetStatus(models.OrderFailed).SetErrorMessage(err.Error()).Save(ctx.Request().Context())
		return c.Fail(err, "failed to complete order")
	}

	_, err = o.Update().SetStatus(models.OrderSuccess).Save(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "failed to update order status")
	}

	log.Printf("complete the order from %s to %s at %s!", o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04"))
	msg.Success(ctx, fmt.Sprintf("complete the order from %s to %s at %s", o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04")))
	return c.Redirect(ctx, "home")
}

func (c *orderController) Delete(ctx echo.Context) error {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		return c.Fail(err, "order id format not correct")
	}

	o, err := c.Container.ORM.Order.Query().Where(order.ID(orderId)).Only(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "order not found")
	}

	orderUser, err := o.QueryUser().Only(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "order have no user")
	}

	currentUser, err := c.Container.Auth.GetAuthenticatedUser(ctx)
	if err != nil {
		return c.Redirect(ctx, "login")
	}

	if orderUser.ID != currentUser.ID {
		return c.Fail(err, "this order is not created by you")
	}

	err = c.Container.ORM.Order.DeleteOneID(orderId).Exec(ctx.Request().Context())
	if err != nil {
		return c.Fail(err, "delete order failed")
	}

	msg.Success(ctx, "order deleted!")
	htmx.SetBoosted(ctx)
	return c.Redirect(ctx, "home")
}
