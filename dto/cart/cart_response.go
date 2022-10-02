package cartdto

type CartResponse struct {
	ID       int `json:"id"`
	BookID   int `json:"book_id"`
	SellerID int `json:"seller_id"`
	UserID   int `json:"user_id"`
	Price    int `json:"price"`
}