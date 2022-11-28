package structs

type Transaction struct {
	DebitedAccountId  uint32 `json:"debitedAccountId"`
	CreditedAccountId uint32 `json:"creditedAccountId"`
	Value             uint64 `json:"value"`
}
