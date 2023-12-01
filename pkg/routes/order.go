package routes

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
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

func (c *orderController) Post(ctx echo.Context) error {
	var form orderForm
	ctx.Set(context.FormKey, &form)

	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse register form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	startTime, err := time.Parse("2006-01-02T15:04", form.StartTime)
	if err != nil {
		form.Submission.SetFieldError("StartTime", "unable to parse start time")
		return c.Get(ctx)
	}
	endTime, err := time.Parse("2006-01-02T15:04", form.EndTime)
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
		SetUserID(user.ID).
		SetStartTime(startTime).
		SetEndTime(endTime).
		SetEmail(form.Email).
		SetPhoneNumber(form.PhoneNumber).
		SetArrivalStation(form.ArrivalStation).
		SetDepartureStation(form.DepartureStation).
		SetIDNumber(form.IdNumber).
		Save(ctx.Request().Context())

	if err != nil {
		// ctx.Logger().Infof("create order failed, err: %s", err)
		log.Printf("create order failed, err: %s", err)
		return c.Fail(err, "unable to create order")
	}

	log.Printf("order from %s to %s at %s is created!", o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04"))
	msg.Success(ctx, fmt.Sprintf("Order from %s to %s at %s is created!", o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04")))
	return c.Redirect(ctx, "order")
}
