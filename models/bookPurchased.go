package models

import "time"

// models that contains books that this user's already purchased!

type BookPurchased struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int    `json:"user_id"`
	BookID        int    `json:"book_id"`
	CreatedAt     time.Time `json:"created_at"` // show when this book is bought
}