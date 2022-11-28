package structs

type Account struct {
	ID         uint32 `json:"id"`
	Balance    uint64 `json:"balance"`
	CustomerId uint32 `json:"customer_id"`
}
