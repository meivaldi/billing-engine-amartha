CREATE TABLE billings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    next_payment_date DATE NOT NULL,
    outstanding INT NOT NULL,
	is_deliquent BOOL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);