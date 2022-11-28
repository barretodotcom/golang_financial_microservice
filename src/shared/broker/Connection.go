package broker

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func GetChannel() *amqp.Channel {

	amqpServerUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", os.Getenv("AMQP_USER"), os.Getenv("AMQP_PASS"), os.Getenv("AMQP_HOST"), os.Getenv("AMQP_PORT"), os.Getenv("AMQP_VHOST"))

	rabbitMqConnection, err := amqp.Dial(amqpServerUrl)

	if err != nil {
		log.Fatalf("Cannot connect to rabbitMQ, reason: %s", err.Error())
	}

	channel, err := rabbitMqConnection.Channel()

	if err != nil {
		log.Fatalf("Cannot get channel to rabbitMQ, reason: %s", err.Error())
	}

	_, err = channel.QueueDeclare(
		"transactions",
		false,
		true,
		false,
		true,
		amqp.Table{
			"x-queue-mode": "lazy",
			"x-queue-type": "classic",
		},
	)

	return channel

}
