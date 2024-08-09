package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
)

const (
	GetBillingQuery = `SELECT id, user_id FROM billings`
	GetPayments     = `SELECT created_at FROM payments WHERE billing_id=$1 ORDER BY created_at DESC`
	SetDeliquent    = `UPDATE billings SET is_deliquent = true WHERE id=$1 AND user_id=$2`
)

type (
	Billing struct {
		BillingID uint64 `json:"billing_id"`
		UserID    uint64 `json:"user_id"`
	}

	Payment struct {
		UserId      uint64    `json:"user_id,omitempty"`
		BillingID   uint64    `json:"billing_id"`
		Amount      uint64    `json:"amount,omitempty"`
		PaymentDate time.Time `json:"payment_date,omitempty"`
	}
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "billingengine", "billing")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to open connection for db, err: %v", err)
		return
	}
	defer db.Close()

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		log.Println("Running cron for deliquent checking")
		rows, err := db.Query(GetBillingQuery)
		if err != nil {
			log.Printf("Error query to db, err: %v\n", err)
			return
		}

		var billings []Billing
		for rows.Next() {
			var billing Billing
			err = rows.Scan(&billing.BillingID, &billing.UserID)
			if err != nil {
				log.Printf("Error read data, err: %v\n", err)
				return
			}
			billings = append(billings, billing)
		}

		for _, bil := range billings {
			rows, err := db.Query(GetPayments, bil.BillingID)
			if err != nil {
				log.Printf("Error query to db, err: %v\n", err)
				return
			}

			var paymentDate []time.Time
			for rows.Next() {
				var date time.Time
				err = rows.Scan(&date)
				if err != nil {
					log.Printf("Error read data, err: %v\n", err)
					return
				}
				paymentDate = append(paymentDate, date)
			}

			count := 0
			for i := 0; i < len(paymentDate)-1; i++ {
				calc := paymentDate[i].Sub(paymentDate[i+1])
				// check whether user is deliquent or not
				if calc > (time.Hour * 168) {
					count++
				}
			}

			if count >= 2 {
				err := db.QueryRow(SetDeliquent, bil.BillingID, bil.UserID)
				if err.Err() != nil {
					log.Printf("Error set deliquent, err: %v\n", err.Err())
					return
				}
			}

			log.Println("Successfully running deliquent checking")
		}
	})

	c.Start()

	select {}
}
