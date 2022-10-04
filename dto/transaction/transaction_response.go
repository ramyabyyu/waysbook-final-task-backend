package transactiondto

import "waysbook/models"

type TransactionResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	User   models.User `json:"user"`
	SellerID int `json:"seller_id"`
	BookID int `json:"book_id"`
	BookPurchased models.Book `json:"book_purchased"`
	Total int `json:"total"`
	Status string `json:"status"`
}