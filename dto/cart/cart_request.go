package cartdto

type CreateCartRequest struct {
	BookID   string `json:"book_id" form:"book_id"`
	SellerID string `json:"seller_id" form:"seller_id"`
	SubTotal string `json:"subtotal" form:"subtotal"`
}

type DeleteCartRequest struct {
	ID     string `json:"cart_id" form:"cart_id"`
	UserID string `json:"user_id" form:"cart_id"`
}