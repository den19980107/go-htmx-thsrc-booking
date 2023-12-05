package models

type OrderStatus string

const (
	OrderPending = "pending"
	OrderWaiting = "waiting"
	OrderSuccess = "success"
	OrderFailed  = "failed"
)
