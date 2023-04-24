package main

import (
	"OrderPayment/service"
	_ "OrderPayment/service"
	"database/sql"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/payment", service.PaymentHandle).Methods("GET")
	http.ListenAndServe(":8081", r)
}
