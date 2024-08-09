package repository

import (
	model "github.com/meivaldi/billing-engine/model"
)

type Repository interface {
	CreateUser(user model.User) (lastId int64, err error)
	GetOutstanding(userId uint64) (model.Billing, error)
	CreateLoan(billing model.Billing) (billingId int64, err error)
	GetDeliquentUsers() ([]model.User, error)
	MakePayment(payment model.Payment) (int64, error)
	MakeRePayment(payment model.Payment) (int64, error)
	IsAlreadyPaid(billingID uint64) (model.Payment, error)
	UpdateBilling(billing model.Billing) error
	GetPaymentData(bilId uint64) ([]model.Payment, error)
	SetDeliquent(bilId uint64) error
}
