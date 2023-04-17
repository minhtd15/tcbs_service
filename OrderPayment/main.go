package main

import (
	"OrderPayment/service"
	_ "OrderPayment/service"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/balance", service.BalanceHandle)
	http.HandleFunc("/payments", service.PaymentHandle)

	fmt.Println("Payment service started on port 9000")
	http.ListenAndServe(":9000", nil)
}
