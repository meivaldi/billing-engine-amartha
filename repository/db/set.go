package db

import (
	"fmt"

	"github.com/meivaldi/billing-engine/model"
)

func (repo *RepositoryDB) CreateUser(user model.User) (lastId int64, err error) {
	var id int64

	err = repo.db.QueryRow(InsertUserQuery, user.Name, user.Age, user.WorkType).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting data, err: %v", err)
	}
	return id, nil
}

func (repo *RepositoryDB) CreateLoan(billing model.Billing) (billingId int64, err error) {
	var id int64

	err = repo.db.QueryRow(CreateBillingQuery, billing.UserID, billing.NextPaymentDate, billing.OutstandingAmount, false).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error create loan, err: %v", err)
	}

	return id, nil
}

func (repo *RepositoryDB) MakePayment(payment model.Payment) (int64, error) {
	var id int64

	err := repo.db.QueryRow(CreatePaymentQuery, payment.BillingID, payment.Amount).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error making payment, err: %v", err)
	}

	return id, nil
}

func (repo *RepositoryDB) MakeRePayment(payment model.Payment) (int64, error) {
	var id int64

	err := repo.db.QueryRow(CreateRePaymentQuery, payment.BillingID, payment.Amount, payment.PaymentDate).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error making payment, err: %v", err)
	}

	return id, nil
}

func (repo *RepositoryDB) UpdateBilling(billing model.Billing) error {
	err := repo.db.QueryRow(UpdateBillingQuery, billing.PaymentAmount, billing.NextPaymentDate, billing.BillingID)
	if err.Err() != nil {
		return fmt.Errorf("error update loan, err: %v", err.Err())
	}

	return nil
}

func (repo *RepositoryDB) SetDeliquent(bilId uint64) error {
	err := repo.db.QueryRow(UpdateDeliquentQuery, bilId)
	if err.Err() != nil {
		return fmt.Errorf("error update loan, err: %v", err.Err())
	}

	return nil
}
