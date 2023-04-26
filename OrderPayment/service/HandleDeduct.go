package service

import (
	"OrderPayment/controller"
	"OrderPayment/entity"
	"encoding/json"
	"net/http"
)

func HandleDeduct(w http.ResponseWriter, r *http.Request) {
	var req entity.DeductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Deduct
	if err := deductBalance(req.UserID, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// return result
	w.WriteHeader(http.StatusOK)
}

func deductBalance(userID int, amount float64) error {
	balance, err := controller.GetBalance(userID)
	if err != nil {
		return UserNotFound
	}
	if balance < amount {
		return InsufficientBalance
	}
	balance -= amount
	return nil
}

var (
	UserNotFound        = &Error{"User not found"}
	InsufficientBalance = &Error{"Insufficient balance"}
)

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}
