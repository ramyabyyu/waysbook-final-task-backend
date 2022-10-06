package transactiondto

type TransactionRequest struct {
	UserID string `json:"user_id"`
	Total  string `json:"total"`
}