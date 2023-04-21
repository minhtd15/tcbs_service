package service

import (
	"OrderPayment/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PaymentHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading payment request:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var paymentReq entity.PaymentRequest
	err = json.Unmarshal(reqBody, &paymentReq)
	if err != nil {
		fmt.Println("Error Unmarshalling payment request:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check balance
	if currentBalance() < paymentReq.Amount {
		fmt.Println("Not enough balance")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Not enough balance"})
		return
	}
	balanceNumber -= paymentReq.Amount
	w.WriteHeader(http.StatusOK)
}
