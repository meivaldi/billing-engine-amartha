package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
	httpDel "github.com/meivaldi/billing-engine/delivery/http"
	dbRepo "github.com/meivaldi/billing-engine/repository/db"
	billingUc "github.com/meivaldi/billing-engine/usecase/billing"
	paymentUc "github.com/meivaldi/billing-engine/usecase/payment"
	userUc "github.com/meivaldi/billing-engine/usecase/user"
)

func main() {
	router := httprouter.New()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "billingengine", "billing")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to open connection for db, err: %v", err)
		return
	}
	defer db.Close()

	repository, err := dbRepo.New(db)
	if err != nil {
		log.Printf("failed to init repository, err: %v", err)
		return
	}

	userUsecase := userUc.New(repository)
	billingUsecase := billingUc.New(repository)
	paymentUsecase := paymentUc.New(repository)

	handler := httpDel.NewHttpHandler(userUsecase, billingUsecase, paymentUsecase)

	// make a loan
	router.POST("/loan/make", handler.MakeLoan)
	router.POST("/loan/pay", handler.MakePayment)
	router.POST("/loan/repay", handler.Repay)

	router.GET("/loan/outstanding/:userId", handler.GetOutstanding)
	router.GET("/loan/deliquent-users", handler.GetDeliquentUsers)

	// Start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", router))
}
