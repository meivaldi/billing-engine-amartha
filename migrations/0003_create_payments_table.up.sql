CREATE TABLE payments(
	id SERIAL PRIMARY KEY,
	billing_id INT NOT NULL,
	amount int NOT NULL,
	created_at DATE NOT NULL
);