package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindBooks() ([]models.Book, error)
	GetBookBySlug(slug string) (models.Book, error) // for book detail
	GetBookByID(ID int) (models.Book, error) // for checking author id in update book thumbnail
	CreateBook(book models.Book) (models.Book, error)
	UpdateBookThumbnail(ID int, bookThumbnail string) (models.Book, error)
	GetUserBook(userID int) ([]models.Book, error) // get books that this user sell
	UpdateBookPromo(ID int, discount int) (models.Book, error) // allow user to make a promo in their book
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

func (r *repository) UpdateBookThumbnail(ID int, bookThumbnail string) (models.Book, error) {
	var book models.Book
	r.db.First(&book, "id=?", ID)

	book.Thumbnail = bookThumbnail
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) GetUserBook(userID int) ([]models.Book, error)  {
	var books []models.Book
	err := r.db.Find(&books, "user_id=?", userID).Error

	return books, err
}

func (r *repository) UpdateBookPromo(ID int, discount int) (models.Book, error) {
	var book models.Book
	r.db.First(&book, "id=?", ID)

	book.IsPromo = true
	book.Discount = discount
	
	// Calculate Price After Discount
	book.PriceAfterDiscount = book.Price - (book.Price * discount / 100)

	err := r.db.Save(&book).Error

	return book, err
}