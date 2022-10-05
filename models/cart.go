package models

import "time"

type Cart struct {
	ID        		int    		`json:"id" gorm:"primary_key:auto_increment"`
	UserID			int			`json:"user_id"`
	User			User		`json:"user"`
	CartItems 		[]CartItem  `json:"cart_items"` // Every cart contain a bunch of cartItem
	IsPay			bool		`json:"is_pay"` // If isPay == true, this means the cart is already been purchased!
	CreatedAt 		time.Time 	`json:"-"`
	UpdatedAt 		time.Time 	`json:"-"`
}

type UserCart struct {
	ID int `json:"id"`
}

func (UserCart) TableName() string {
	return "carts"
}