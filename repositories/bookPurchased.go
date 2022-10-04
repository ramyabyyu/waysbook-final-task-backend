package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookPurchasedRepository interface {
	FindBooksPurcased() ([]models.BookPurchased, error)
	CreateBookPurchased(bookPurchased models.BookPurchased) (models.BookPurchased, error)
}

func RepositoryBookPurchased(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBooksPurcased() ([]models.BookPurchased, error) {
	var bookPurchased []models.BookPurchased
	err := r.db.Find(&bookPurchased).Error

	return bookPurchased, err
}

func (r *repository) CreateBookPurchased(bookPurchased models.BookPurchased) (models.BookPurchased, error) {
	err := r.db.Create(&bookPurchased).Error

	return bookPurchased, err
}