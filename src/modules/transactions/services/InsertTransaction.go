package transaction_services

import (
	"transactions_microservice/src/shared/database"
)

func InsertTransaction(debitedAccountId uint32, creditedAccountId uint32, value uint64) {

	db := database.Connection()

	db.Query("INSERT INTO transactions_dbs(debited_account_id, credited_account_id,value) VALUES($1, $2, $3)", debitedAccountId, creditedAccountId, value)
}
