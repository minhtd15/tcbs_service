package client

import (
	"OrderService/entity"
	"encoding/json"
	"github.com/streadway/amqp"
	"net/http"
)

func RabbitSender(w http.ResponseWriter, order entity.OrderRequest) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"order_exchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ch.Publish(
		"order_exchange",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
