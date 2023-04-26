package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func DeductBalance(userID int, amount float64) error {
	reqBody, err := json.Marshal(map[string]interface{}{
		"user_id": userID,
		"amount":  amount,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8081/payment/deduct", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check deduct balance from service payment
	if resp.StatusCode != http.StatusOK {
		return errors.New("Deduct balance failed")
	}
	return nil
}
