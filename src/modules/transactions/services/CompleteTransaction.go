package transaction_services

import (
	"transactions_microservice/src/utils/structs"
)

func CompleteTransaction(debitedAccount structs.Account, creditedAccount structs.Account, value uint64) {
	InsertTransaction(debitedAccount.ID, creditedAccount.ID, value)

	go CalcAccountBalance(debitedAccount.ID)
	go CalcAccountBalance(creditedAccount.ID)
}
