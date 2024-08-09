package model

import "time"

type Billing struct {
	BillingID         uint64    `json:"billing_id"`
	UserID            uint64    `json:"user_id"`
	OutstandingAmount uint64    `json:"outstanding_amount"`
	NextPaymentDate   time.Time `json:"next_payment_date"`
	PaymentAmount     uint64    `json:"payment_amount,omitempty"`
	IsDeliquent       bool      `json:"is_deliquent"`
}
