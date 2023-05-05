package main

import (
	"OrderPayment/client"
	"OrderPayment/entity"
	"OrderPayment/service"
	_ "OrderPayment/service"
	"database/sql"
	"encoding/json"
	_ "github.com/gorilla/mux"
	"log"
	_ "net/http"
)

var db *sql.DB

func main() {
	// process message, update the balance
	//r := mux.NewRouter()
	//r.HandleFunc("/payment/balance/", service.HandleBalance)
	//r.HandleFunc("/payment/deduct", service.HandleDeduct)
	//log.Fatal(http.ListenAndServe(":8081", r))
	//log.Printf("Payment completed")

	msgs := client.RabbitConsumer()

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var req entity.DeductRequest
			err := json.Unmarshal(d.Body, &req)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			// Deduct Balance from user's account
			err = service.HandleDeduct(req.UserID, req.Amount)

			// send a success message to RabbitMQ queue
			successResp := entity.PaymentResponse{
				Success: true,
				Message: "Payment processed successfully",
			}

			successRespBytes, err := json.Marshal(successResp)
			if err != nil {
				log.Printf("Failed to marshal response: %v", err)
				continue
			}
			err = client.SendMessage(successRespBytes)
			if err != nil {
				log.Printf("Failed to send success message: %v", err)
				continue
			}
		}
	}()

	log.Printf("Waiting for messages")
	<-forever
}
