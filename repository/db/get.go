package db

import (
	"github.com/meivaldi/billing-engine/model"
)

func (repo *RepositoryDB) GetOutstanding(userId uint64) (model.Billing, error) {
	rows, err := repo.db.Query(GetUserQuery, userId)
	if err != nil {
		return model.Billing{}, err
	}
	defer rows.Close()

	var billing model.Billing
	for rows.Next() {
		err = rows.Scan(&billing.BillingID, &billing.UserID, &billing.OutstandingAmount, &billing.NextPaymentDate, &billing.IsDeliquent)
		if err != nil {
			return model.Billing{}, err
		}
	}

	return billing, nil
}

func (repo *RepositoryDB) GetDeliquentUsers() ([]model.User, error) {
	rows, err := repo.db.Query(GetDeliquentUsersQuery)
	if err != nil {
		return []model.User{}, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.UserID, &user.Name, &user.Age, &user.WorkType)
		if err != nil {
			return []model.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *RepositoryDB) IsAlreadyPaid(billingID uint64) (model.Payment, error) {
	rows, err := repo.db.Query(SelectPaymentQuery, billingID)
	if err != nil {
		return model.Payment{}, err
	}

	var payment model.Payment
	for rows.Next() {
		err = rows.Scan(&payment.Id, &payment.BillingID, &payment.Amount, &payment.PaymentDate)
		if err != nil {
			return model.Payment{}, err
		}
	}

	return payment, nil
}

func (repo *RepositoryDB) GetPaymentData(bilId uint64) ([]model.Payment, error) {
	rows, err := repo.db.Query(GetPaymentDataQuery, bilId)
	if err != nil {
		return []model.Payment{}, err
	}
	defer rows.Close()

	var payments []model.Payment
	for rows.Next() {
		var payment model.Payment
		err = rows.Scan(&payment.PaymentDate)
		if err != nil {
			return []model.Payment{}, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}
