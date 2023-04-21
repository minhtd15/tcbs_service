package main

import (
	"OrderPayment/service"
	_ "OrderPayment/service"
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

func main() {
	http.HandleFunc("/balance/{user_id}", service.BalanceHandle)
	http.HandleFunc("/payments", service.PaymentHandle)

	fmt.Println("Payment service started on port 9000")
	http.ListenAndServe(":9000", nil)
}
