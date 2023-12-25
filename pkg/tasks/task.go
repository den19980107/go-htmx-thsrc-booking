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

const Hourly = "0 * * * *"

func Run(c *services.Container) {
	cron := cronV3.New()
	cron.Start()

	booking(c, context.Background())
	cron.AddJob(Hourly, cronV3.FuncJob(func() {
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

	loc, _ := time.LoadLocation("Asia/Taipei")
	today := time.Now().In(loc)
	avaliableStartTime := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, loc)
	avaliableEndTime := avaliableStartTime.Add(29 * 24 * time.Hour)
	log.Printf("today is %s, retrive orders which between %s and %s", today, avaliableStartTime, avaliableEndTime)

	orders, err := c.ORM.Order.Query().Where(order.StartTimeGTE(avaliableStartTime), order.StartTimeLT(avaliableEndTime), order.StatusEQ(models.OrderPending)).All(context.Background())
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
