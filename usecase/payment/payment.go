package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/meivaldi/billing-engine/model"
	"github.com/meivaldi/billing-engine/repository"

	uc "github.com/meivaldi/billing-engine/usecase"
)

const (
	Amount = 110000
)

type PaymentUsecase struct {
	repositoryDB repository.Repository
}

func New(repo repository.Repository) uc.IPaymentUsecase {
	return &PaymentUsecase{
		repositoryDB: repo,
	}
}

func (uc *PaymentUsecase) MakePayment(payment model.Payment) (int64, error) {
	payData, err := uc.repositoryDB.IsAlreadyPaid(payment.BillingID)
	if err != nil {
		return 0, err
	}

	if payment.Amount < Amount {
		return 0, errors.New("insufficient amount")
	}

	duration := time.Since(payData.PaymentDate)
	if duration < time.Hour*168 {
		return 0, errors.New("already paid")
	}

	paymentId, err := uc.repositoryDB.MakePayment(payment)
	if err != nil {
		return 0, err
	}

	billingData := model.Billing{
		BillingID:       payment.BillingID,
		PaymentAmount:   Amount,
		NextPaymentDate: time.Now().Add(168 * time.Hour),
	}

	err = uc.repositoryDB.UpdateBilling(billingData)
	if err != nil {
		return 0, err
	}

	return paymentId, nil
}

func (uc *PaymentUsecase) Repay(payment model.Payment) ([]int64, error) {
	paymentData, err := uc.repositoryDB.GetPaymentData(payment.BillingID)
	if err != nil {
		return []int64{}, err
	}

	payload := []model.Payment{}
	totalAmount := 0
	for i := 0; i < len(paymentData)-1; i++ {
		calc := paymentData[i].PaymentDate.Sub(paymentData[i+1].PaymentDate)
		// check whether user is deliquent or not
		if calc > (time.Hour * 168) {
			payDate := paymentData[i+1].PaymentDate.AddDate(0, 0, 7)
			pay := model.Payment{
				BillingID:   payment.BillingID,
				Amount:      Amount,
				PaymentDate: payDate,
			}
			totalAmount += Amount
			payload = append(payload, pay)
		}
	}

	if len(payload) == 0 {
		return []int64{}, errors.New("payment is on schedule, no deliquent")
	}

	if payment.Amount < uint64(totalAmount) {
		return []int64{}, errors.New("insufficient amount")
	}

	ids := []int64{}
	isRepaySuccess := true

	for _, data := range payload {
		id, err := uc.repositoryDB.MakeRePayment(data)
		if err != nil {
			isRepaySuccess = false
			log.Print(err)
			continue
		}
		ids = append(ids, id)
	}

	if isRepaySuccess {
		err = uc.repositoryDB.SetDeliquent(payment.BillingID)
		if err != nil {
			return []int64{}, err
		}
	}

	return ids, nil
}
