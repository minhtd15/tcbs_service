package main

import (
	"OrderPayment/service"
	_ "OrderPayment/service"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	//http.HandleFunc("/balance/{userID}", service.BalanceHandle)
	//http.HandleFunc("/payments", service.PaymentHandle)
	//
	//fmt.Println("Payment service started on port 9000")
	//http.ListenAndServe(":9000", nil)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ server: %v", err)
	}
	defer conn.Close()

	// create channel to send and receive message
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open the channel: %v", err)
	}
	defer conn.Close()

	// declare queue to receive message from service order
	q, err := ch.QueueDeclare(
		"payment_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// register consumer for queue and process message when received
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// process message, update the balance
			r := mux.NewRouter()
			//r.HandleFunc("/payment/balance/", service.HandleBalance)
			r.HandleFunc("/payment/deduct", service.HandleDeduct)
			log.Fatal(http.ListenAndServe(":8081", r))
			log.Printf("Payment completed")
		}
	}()

	log.Printf("Waiting for messages")
	<-forever

}
