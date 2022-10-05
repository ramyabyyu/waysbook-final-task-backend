package models

type CartItem struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	CartID        int    `json:"cart_id"`
	Cart          Cart   `json:"cart"`
	UserID        int    `json:"user_id"`
	BuyerName     string `json:"buyer_name"`
	BookID        int    `json:"book_id"` // book purchased
	BookTitle     string `json:"book_title"`
	BookThumbnail string `json:"book_thumbnail"`
	SellerID      int    `json:"seller_id"` // get this data from book
	SellerName    string `json:"seller_name"`
	Qty           int    `json:"qty"`   // default 1, for now
	Price         int    `json:"price"` // get this data from book_id
}

type CartItemResponse struct {
	ID int `json:"id"`
}

func (CartItemResponse) TableName() string {
	return "cart_items"
}