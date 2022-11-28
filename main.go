package main

import (
	"log"
	transaction_services "transactions_microservice/src/modules/transactions/services"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

	transaction_services.ConsumeMessages()

}
