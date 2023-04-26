package service

import (
	"OrderPayment/controller"
	"encoding/json"
	"net/http"
	"strconv"
)

var payments = make(map[string]float64)

func HandleBalance(w http.ResponseWriter, r *http.Request) {
	// get userID from requestURL
	userID := r.URL.Path[len("/payment/balance/"):]

	tmp, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get the balance from customer's account
	balance, err := controller.GetBalance(tmp)

	// return balance
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}
