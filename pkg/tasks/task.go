package tasks

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/pkg/models"
	"github.com/mikestefanello/pagoda/pkg/services"
	cronV3 "github.com/robfig/cron/v3"
)

func Run(c *services.Container) {
	cron := cronV3.New()
	cron.Start()

	booking(c, context.Background())
	cron.AddJob(fmt.Sprintf("@every %s", time.Minute), cronV3.FuncJob(func() {
		booking(c, context.Background())
	}))
}

func booking(c *services.Container, ctx context.Context) {
	log.Printf("================ start booking jobs ... ==================")
	var jobError error
	defer func() {
		if jobError != nil {
			log.Printf("complete booking job with err: %s", jobError)
		} else {
			log.Printf("complete booking job!")
		}
		log.Println()
	}()

	log.Println("retrive orders which need to book ...")

	today := time.Now().UTC()
	orders, err := c.ORM.Order.Query().Where(order.StartTimeGT(today), order.StatusEQ(models.OrderPending)).All(context.Background())
	if err != nil {
		jobError = fmt.Errorf("get orders failed, err: %s", err)
		return
	}
	log.Printf("start booking %d orders ...", len(orders))

	errs := []error{}
	for _, order := range orders {
		log.Printf("start booking order id: %d", order.ID)
		err := Booking(context.Background(), c, order)
		if err != nil {
			errs = append(errs, fmt.Errorf("booking order %d failed, err: %s", order.ID, err))
		}
	}

	errsStr := ""
	for _, err := range errs {
		errsStr += err.Error()
	}

	if errsStr != "" {
		jobError = errors.New(errsStr)
	}
}
