package models

import "time"

type Book struct {
	ID              int    `json:"id" gorm:"primary_key:auto_increment"`
	Title           string `json:"title" gorm:"type: varchar(255)"`
	Slug string `json:"slug"`
	PublicationDate time.Time `json:"publication_date"`
	Pages int `json:"pages"`
	ISBN int `json:"ISBN"`
	UserID int `json:"user_id"`
	User User `json:"user"`
	Price int `json:"price"`
	Description string `json:"description" gorm:"type: text"`
	BookAttachment string `json:"book_attachment"`
	Thumbnail string `json:"thumbnail"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserBook struct {
	ID int `json:"id"`
}

func (UserBook) TableName() string {
	return "books"
}