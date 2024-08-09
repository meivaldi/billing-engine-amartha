package db

const (
	InsertUserQuery = `
		INSERT INTO users (name, age, work_type) 
		VALUES ($1, $2, $3)
		RETURNING id
	`

	CreateBillingQuery = `
		INSERT INTO billings(user_id, next_payment_date, outstanding, is_deliquent, updated_at) 
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
		RETURNING id
	`

	CreatePaymentQuery = `
		INSERT INTO payments(billing_id, amount, created_at) 
		VALUES ($1, $2, CURRENT_TIMESTAMP)
		RETURNING id
	`

	CreateRePaymentQuery = `
		INSERT INTO payments(billing_id, amount, created_at) 
		VALUES ($1, $2, $3)
		RETURNING id
	`

	GetUserQuery = `
		SELECT id, user_id, outstanding, next_payment_date, is_deliquent FROM billings WHERE user_id = $1
	`

	GetDeliquentUsersQuery = `
		SELECT 
			users.id AS user_id,
			users.name,
			users.age,
			users.work_type
		FROM 
			users
		INNER JOIN
			billings ON users.id = billings.user_id
		WHERE 
			is_deliquent = true
	`

	UpdateBillingQuery = `
		UPDATE
			billings
		SET
			outstanding = outstanding-$1,
			next_payment_date = $2
		WHERE
			id = $3;
	`

	UpdateDeliquentQuery = `UPDATE billings SET is_deliquent = false WHERE id=$1`

	SelectPaymentQuery = `
		SELECT * FROM payments WHERE billing_id=$1 ORDER BY created_at DESC LIMIT 1
	`

	GetPaymentDataQuery = `
		SELECT created_at FROM payments WHERE billing_id=$1 ORDER BY created_at DESC
	`
)
