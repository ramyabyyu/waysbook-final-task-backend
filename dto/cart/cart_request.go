package cartdto

type CreateCartItemRequest struct {
	BookID string `json:"book_id" form:"book_id"`
	CartID string `json:"cart_id" form:"cart_id"`
}

type GetCartItemRequest struct {
	CartID string `json:"cart_id"`
}

type DeleteCartRequest struct {
	ID     string `json:"cart_id" form:"cart_id"`
	UserID string `json:"user_id" form:"cart_id"`
}