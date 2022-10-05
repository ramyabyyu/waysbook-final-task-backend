package transactiondto

type TransactionRequest struct {
	CartID string `json:"cart_id"`
	Total  string `json:"total"`
}