package models

import "time"

type Transaction struct {
	ID        int     `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int    `json:"user_id"`
	User User `json:"user"`
	SellerID int `json:"seller_id"`
	Carts []Cart `json:"cart"`
	Total int `json:"total"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

