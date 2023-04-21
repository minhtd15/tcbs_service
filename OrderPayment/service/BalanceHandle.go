package service

import (
	"OrderPayment/controller"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var balanceNumber float64

func BalanceHandle(w http.ResponseWriter, r *http.Request) {
	// read the request in json
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	balanceNumber, err = controller.GetBalance(userID)

	//balanceNumber, err := json.Marshal(entity.Balance{currentBalance()})
	if err != nil {
		fmt.Println("Error marshal data:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balanceNumber)
}

func currentBalance() float64 {
	return balanceNumber
}
