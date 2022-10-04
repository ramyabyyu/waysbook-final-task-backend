package models

type BookPurchased struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int    `json:"user_id"`
	BookID        int    `json:"book_id"`
	BookTitle     string `json:"book_title"`
	BookThumbnail string `json:"book_thumbnail"`
}