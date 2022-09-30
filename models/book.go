package models

import "time"

type Book struct {
	ID              int    `json:"id" gorm:"primary_key:auto_increment"`
	Title           string `json:"title" gorm:"type: varchar(255)"`
	Slug string `json:"slug"`
	PublicationDate time.Time `json:"publication_date"`
	Pages int `json:"pages"`
	ISBN int `json:"ISBN"`
	Price int `json:"price"`
	IsPromo bool `json:"is_promo"`
	Discount int `json:"discount"` //* -> e.g dicount = 45, it means 45%
	PriceAfterDiscount int `json:"price_after_discount"` //* result from the old price after getting discount promo
	Description string `json:"description" gorm:"type: text"`
	BookAttachment string `json:"book_attachment"`
	Thumbnail string `json:"thumbnail"`
	UserID int `json:"user_id"`
	User User `json:"user"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserBook struct {
	ID int `json:"id"`
}

func (UserBook) TableName() string {
	return "books"
}