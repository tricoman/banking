package dto

type NewTransactionResponse struct {
	TransactionId  string  `json:"transaction_id"`
	UpdatedBalance float64 `json:"updated_balance"`
}
