package transaction_services

import "transactions_microservice/src/shared/database"

func CalcAccountBalance(accountId uint32) {
	db := database.Connection()

	var balance uint64

	db.QueryRow(`
		SELECT
		COALESCE(SUM(t.value), 0)+ 10000 - (
			SELECT
				COALESCE(SUM(t2.value), 0)
			FROM
				transactions_dbs t2
			WHERE
				t2."debited_account_id" = $1
		) 
		AS "balance"
		FROM
			transactions_dbs t
		WHERE
		t."credited_account_id" = $1
	`, accountId).Scan(&balance)

	db.QueryRow(`
		UPDATE
			account_dbs
		SET balance = $1
		WHERE id = $2
	`, balance, accountId)

}
