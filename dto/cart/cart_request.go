package cartdto

type CartRequest struct {
	BookID   string `json:"book_id" form:"book_id"`
	SellerID string `json:"seller_id" form:"seller_id"`
	Price    string `json:"price" form:"price"`
}