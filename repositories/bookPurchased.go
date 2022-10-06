package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookPurchasedRepository interface {
	GetOneBook(ID int) (models.Book, error)
	FindBooksPurcased(userID int) ([]models.BookPurchased, error)
	CreateBookPurchased(bookPurchased models.BookPurchased)
}

func RepositoryBookPurchased(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetOneBook(ID int) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *repository) CreateBookPurchased(bookPurchased models.BookPurchased) {
	r.db.Create(&bookPurchased)
}

func (r *repository) FindBooksPurcased(userID int) ([]models.BookPurchased, error) {
	var bookPurchased []models.BookPurchased
	err := r.db.Find(&bookPurchased, "user_id=?", userID).Error

	return bookPurchased, err
}