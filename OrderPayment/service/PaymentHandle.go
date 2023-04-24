package service

import (
	"OrderPayment/controller"
	"fmt"
	"net/http"
	"strconv"
)

func PaymentHandle(w http.ResponseWriter, r *http.Request) {
	// receive userID from service oreder
	userID := r.URL.Query().Get("user_id")
	tmp, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// receive amount from

	balance := controller.GetBalance(tmp)

	// check the valid balance
	if amount < 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if amount < balance {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "Insufficient balance"}`)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "OK"}`)

	//reqBody, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Println("Error reading payment request:", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//var paymentReq entity.PaymentRequest
	//err = json.Unmarshal(reqBody, &paymentReq)
	//if err != nil {
	//	fmt.Println("Error Unmarshalling payment request:", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//// check balance
	//if currentBalance() < paymentReq.Amount {
	//	fmt.Println("Not enough balance")
	//	w.WriteHeader(http.StatusBadRequest)
	//	json.NewEncoder(w).Encode(map[string]string{"error": "Not enough balance"})
	//	return
	//}
	//balance -= paymentReq.Amount
	//w.WriteHeader(http.StatusOK)
}
