package cartdto

type CartResponse struct {
	ID            int    `json:"id"`
	QTY           int    `json:"qty"`
	Subtotal      int    `json:"subtotal"`
	SellerID      int    `json:"seller_id"`
	BookID        int    `json:"book_id"`
	UserID        int    `json:"user_id"`
	BookTitle     string `json:"book_title"`
	BookThumbnail string `json:"book_thumbnail"`
	Author        string `json:"author"`
	Slug          string `json:"slug"`
}

type DeleteCartResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}