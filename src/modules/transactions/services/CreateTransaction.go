package transaction_services

import (
	"encoding/json"
	"fmt"
	account_services "transactions_microservice/src/modules/account/services"
	"transactions_microservice/src/utils/structs"
)

func CreateTransaction(message []byte) {
	var transaction structs.Transaction

	err := json.Unmarshal(message, &transaction)

	if err != nil {
		fmt.Print("Message is not a json", err.Error())
	}
	debitedAccount := make(chan structs.Account)
	creditedAccount := make(chan structs.Account)

	go account_services.FindAccountById(transaction.DebitedAccountId, debitedAccount)
	go account_services.FindAccountById(transaction.CreditedAccountId, creditedAccount)

	CompleteTransaction(<-debitedAccount, <-creditedAccount, transaction.Value)
}
