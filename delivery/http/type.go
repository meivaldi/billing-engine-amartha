package http

import (
	"time"

	"github.com/meivaldi/billing-engine/model"
)

type HttpResponse struct {
	Err     bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseUserData struct {
	model.User
	Amount          uint64    `json:"amount,omitempty"`
	NextPaymentDate time.Time `json:"next_payment_date,omitempty"`
}
