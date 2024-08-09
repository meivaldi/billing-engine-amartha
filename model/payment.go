package model

import "time"

type Payment struct {
	Id          uint64    `json:"payment_id,omitempty"`
	UserId      uint64    `json:"user_id,omitempty"`
	BillingID   uint64    `json:"billing_id"`
	Amount      uint64    `json:"amount,omitempty"`
	PaymentDate time.Time `json:"payment_date,omitempty"`
}
