package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	GetBookBySlug(slug string) (models.Book, error) // for book detail
	GetBookByID(ID int) (models.Book, error) // for checking author id in update book attachment and thumbnail
	CreateBook(book models.Book) (models.Book, error)
	UpdateBookAttachment(ID int, bookAttachment string) (models.Book, error)
	UpdateBookThumbnail(ID int, bookThumbnail string) (models.Book, error)
}

func RepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) GetBookBySlug(slug string) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "slug=?", slug).Error

	return book, err
}

func (r *repository) GetBookByID(ID int) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("User").First(&book, "id=?", ID).Error

	return book, err
}

func (r *repository) CreateBook(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) UpdateBookAttachment(ID int, bookAttachment string) (models.Book, error) {
	var book models.Book
	r.db.First(&book, "id=?", ID)

	book.BookAttachment = bookAttachment
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) UpdateBookThumbnail(ID int, bookThumbnail string) (models.Book, error) {
	var book models.Book
	r.db.First(&book, "id=?", ID)

	book.Thumbnail = bookThumbnail
	err := r.db.Save(&book).Error

	return book, err
}