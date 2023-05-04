package service

import (
	"OrderService/client"
	"OrderService/entity"
	_ "bytes"
	"encoding/json"
	"net/http"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	order := entity.OrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//// call service payment to check the balance of customers
	//resp, err := client.OrderClient(order, w)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//// read the balance from response body of service payment
	//var balance float64
	//if err := json.NewDecoder(resp.Body).Decode(&balance); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//// check the balance of the account and deduct if it's enough
	//if balance < order.Amount {
	//	http.Error(w, "Insufficient balance", http.StatusBadRequest)
	//	return
	//}
	//// Deduct balance
	//if err := DeductBalance(order.UserID, order.Amount); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Send message to RabbitMQ
	orderBytes, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.RabbitSender(orderBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Trả về thông báo đơn hàng thành công
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order successful"))
}
