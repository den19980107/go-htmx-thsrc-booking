package tasks

import (
	"context"
	"fmt"
	"log"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/pkg/models"
	"github.com/mikestefanello/pagoda/pkg/services"
)

func Booking(ctx context.Context, c *services.Container, o *ent.Order) error {
	err := c.CrawlerStore.Create(o.ID)
	if err != nil {
		return fmt.Errorf("create crawler failed, err: %s", err)
	}
	cw, _ := c.CrawlerStore.Get(o.ID)

	tx, err := c.ORM.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start trasaction, err: %s", err)
	}
	defer tx.Commit()

	log.Printf("order user %s booking from %s to %s at %s", o.Email, o.DepartureStation, o.ArrivalStation, o.StartTime.Format("2006-01-02 15:04"))
	imageBase64, jsessiongId, cookieStr, err := cw.GetCaptchaImageAndJsessionId(o.ID)

	if err != nil {
		_, err := tx.Order.Update().Where(order.ID(o.ID)).SetStatus(models.OrderFailed).Save(ctx)
		if err != nil {
			return fmt.Errorf("update order %d from processing to failed failed", o.ID)
		}
		return fmt.Errorf("order %d process failed, err: %s", o.ID, err)
	}

	orderValidation, err := tx.OrderValidation.Create().SetOrder(o).SetJessionID(jsessiongId).SetCaptchaImage(imageBase64).SetCookies(cookieStr).Save(ctx)
	if err != nil {
		return fmt.Errorf("save order validation to db failed, err: %s", err)
	}

	_, err = tx.Order.Update().Where(order.ID(o.ID)).SetStatus(models.OrderWaiting).Save(ctx)
	if err != nil {
		return fmt.Errorf("update order %d from processing to success failed", o.ID)
	}

	err = c.Mail.
		Compose().
		To(o.Email).
		Subject("Please enter the verification code").
		Body(fmt.Sprintf("Go here to enter your verification code: http://%s:%d/order/validate/%d", c.Config.HTTP.Domain, c.Config.HTTP.Port, orderValidation.ID)).
		Send()

	if err != nil {
		return fmt.Errorf("send email to user failed, err: %s", err)
	}

	log.Printf("order %d process success", o.ID)
	return nil
}
