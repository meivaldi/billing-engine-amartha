package usecase

import (
	"errors"

	"github.com/meivaldi/billing-engine/model"
	"github.com/meivaldi/billing-engine/repository"
	uc "github.com/meivaldi/billing-engine/usecase"
)

type BillingUsecase struct {
	repositoryDB repository.Repository
}

func New(dbRepo repository.Repository) uc.IBillingUsecase {
	return &BillingUsecase{
		repositoryDB: dbRepo,
	}
}

func (uc *BillingUsecase) CreateLoan(billing model.Billing) (int64, error) {
	id, err := uc.repositoryDB.CreateLoan(billing)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *BillingUsecase) GetOutstanding(userId uint64) (model.Billing, error) {
	billing, err := uc.repositoryDB.GetOutstanding(userId)
	if err != nil {
		return model.Billing{}, err
	}

	if billing.BillingID == 0 {
		return model.Billing{}, errors.New("billing not found")
	}

	return billing, nil
}
