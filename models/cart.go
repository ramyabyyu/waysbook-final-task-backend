package models

import "time"

type Cart struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	BookID    int 	`json:"book_id"`
	Price int `json:"price"` // price of the product
	SellerID int `json:"seller_id"` // seller ID / user that sell the products
	UserID    int    `json:"user_id"` // buyer ID / user that owns this cart
	User      User   `json:"user"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserCart struct {
	ID int `json:"id"`
}

func (UserCart) TableName() string {
	return "carts"
}