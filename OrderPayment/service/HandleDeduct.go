package service

import (
	"OrderPayment/controller"
	"fmt"
	"log"
)

func HandleDeduct(userID int, amount float64) error {
	//var req entity.DeductRequest
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	// Deduct
	if err := deductBalance(userID, amount); err != nil {
		return fmt.Errorf("Failed to deduct and update balance: %v", err)
	}
	// return result
	log.Printf("Deducted %v from account %v", amount, userID)
	return nil
}

func deductBalance(userID int, amount float64) error {
	fmt.Println("da tru")
	balance, err := controller.GetBalance(userID)
	fmt.Printf("da tru %v and %v \n", amount, balance)
	if err != nil {
		return UserNotFound
	}

	if balance < amount {
		return InsufficientBalance
	}
	balance -= amount
	fmt.Println(balance)

	balance, err = controller.UpdateBalance(balance, userID)
	if err != nil {
		return err
	}

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
