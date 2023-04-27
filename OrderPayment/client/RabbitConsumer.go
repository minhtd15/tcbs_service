package client

import (
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

func RabbitConsumer(w http.ResponseWriter) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ server: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"payment_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to decalre a queue: %v", err)
	}

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
		log.Fatalf("Falied to register a consumer: %v", err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %v", d.Body)

			// xu ly message, cap nhat tai khoan cua khach hang
			// ...

			log.Printf("Payment completed")

		}
	}()

	log.Printf("Waiting for messages...")
	<-forever
}
