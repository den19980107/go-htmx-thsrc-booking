package routes

import (
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/station"

	"github.com/labstack/echo/v4"
)

type (
	home struct {
		controller.Controller
	}

	order struct {
		Id               int
		StartTime        string
		EndTime          string
		DepartureStation station.StationCode
		ArrivalStation   station.StationCode
		IdNumber         string
		PhoneNumber      string
		Email            string
	}
)

func (c *home) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "home"
	page.Metatags.Description = "Welcome to the homepage."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}
	page.Pager = controller.NewPager(ctx, 4)
	orders, err := c.getUserOrders(ctx, &page.Pager)
	if err != nil {
		return c.Fail(err, "failed to get user orders")
	}
	page.Data = orders

	return c.RenderPage(ctx, page)
}

func (c *home) getUserOrders(ctx echo.Context, pager *controller.Pager) ([]order, error) {
	user, err := c.Container.Auth.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	orders, err := user.QueryOrderer().Limit(pager.ItemsPerPage).Offset(pager.GetOffset()).All(ctx.Request().Context())
	if err != nil {
		return nil, err
	}

	result := []order{}
	for _, o := range orders {
		result = append(result, order{
			Id:               o.ID,
			StartTime:        o.StartTime.Format("2006-01-02 15:04"),
			EndTime:          o.EndTime.Format("2006-01-02 15:04"),
			DepartureStation: station.StationCode(o.DepartureStation),
			ArrivalStation:   station.StationCode(o.ArrivalStation),
			IdNumber:         o.IDNumber,
			PhoneNumber:      o.PhoneNumber,
			Email:            o.Email,
		})
	}

	return result, nil
}
