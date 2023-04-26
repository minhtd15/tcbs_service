package main

import (
	"OrderPayment/service"
	_ "OrderPayment/service"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	//http.HandleFunc("/balance/{userID}", service.BalanceHandle)
	//http.HandleFunc("/payments", service.PaymentHandle)
	//
	//fmt.Println("Payment service started on port 9000")
	//http.ListenAndServe(":9000", nil)

	r := mux.NewRouter()
	//r.HandleFunc("/payment/balance/", service.HandleBalance)
	r.HandleFunc("/payment/deduct", service.HandleDeduct)
	log.Fatal(http.ListenAndServe(":8081", r))
}
