package transaction_services

import (
	"fmt"
	"transactions_microservice/src/shared/broker"

	"github.com/streadway/amqp"
)

func ConsumeMessages() {
	channel := broker.GetChannel()

	messages, err := channel.Consume("transactions", "", false, false, false, false, amqp.Table{"x-queue-mode": "lazy"})

	if err != nil {
		fmt.Println(err.Error())
	}

	for message := range messages {
		go CreateTransaction(message.Body)
	}
}
