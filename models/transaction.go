package models

import "time"

type Transaction struct {
	ID        	int     	`json:"id" gorm:"primary_key:auto_increment"`
	Total 		int 		`json:"total"`
	CartID		int			`json:"cart_id"`
	Cart		Cart		`json:"cart"`
	Status 		string 		`json:"status"`
	CreatedAt 	time.Time 	`json:"-"`
	UpdatedAt 	time.Time 	`json:"-"`
}

type CartTransaction struct {
	ID int `json:"id"`
}

func (CartTransaction) TableName() string {
	return "transactions"
}