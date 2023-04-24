package service

import (
	"OrderService/entity"
	_ "bytes"
	"encoding/json"
	"net/http"
)

func OrderService(w http.ResponseWriter, r *http.Request) {
	order := entity.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/payment", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	q := req.URL.Query()
	q.Add("user_id", order.UserID)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var paymentResponse entity.PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// response for orderHandler
	if !paymentResponse.Status {
		http.Error(w, "Transaction failed. Invalid balance", http.StatusBadRequest)
	}
	// encode response body as JSON and send response
	w.WriteHeader(http.StatusOK)
}
