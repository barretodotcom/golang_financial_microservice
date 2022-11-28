package account_services

import (
	"fmt"
	"transactions_microservice/src/shared/database"
	"transactions_microservice/src/utils/structs"
)

func FindAccountById(accountId uint32, account chan structs.Account) {
	db := database.Connection()

	var accountModel structs.Account
	var id uint32
	var customer_id uint32
	var balance uint64
	err := db.QueryRow("SELECT id, customer_id, balance FROM account_dbs WHERE id = $1", accountId).Scan(&id, &customer_id, &balance)

	if err != nil {
		fmt.Println(err)
	}

	accountModel.ID = id
	accountModel.CustomerId = customer_id
	accountModel.Balance = balance

	account <- accountModel
}
