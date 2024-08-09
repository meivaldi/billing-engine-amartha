package usecase

import (
	model "github.com/meivaldi/billing-engine/model"
)

type IUserUsecase interface {
	CreateUser(user model.User) (int64, error)
	GetDeliquentUsers() ([]model.User, error)
}

type IBillingUsecase interface {
	CreateLoan(billing model.Billing) (int64, error)
	GetOutstanding(userId uint64) (model.Billing, error)
}

type IPaymentUsecase interface {
	MakePayment(payment model.Payment) (int64, error)
	Repay(payment model.Payment) ([]int64, error)
}
