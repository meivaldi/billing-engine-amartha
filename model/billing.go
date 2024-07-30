package model

import "time"

type Billing struct {
	userID          uint64    `json:"user_id"`
	nextPaymentDate time.Time `json:"next_payment_date"`
	isDeliquent     bool      `json:"isDeliquent"`
}
