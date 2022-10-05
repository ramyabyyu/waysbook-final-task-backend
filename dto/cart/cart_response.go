package cartdto

type CartItemResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`   // buyer
	SellerID      int    `json:"seller_id"` // seller
	BookID        int    `json:"book_id"`
	BuyerName     string `json:"buyer_name"`
	SellerName    string `json:"seller_name"`
	BookTitle     string `json:"book_title"`
	BookThumbnail string `json:"book_thumbnail"`
	Price         int    `json:"price"`
	Qty           int    `json:"qty"`
}
type DeleteCartResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}