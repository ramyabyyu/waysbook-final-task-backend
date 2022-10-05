package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCartItems(cartID int) ([]models.CartItem, error)
	GetBook(bookID int) (models.Book, error)
	GetUser(userID int) (models.User, error)
	CreateCartItem(cartItem models.CartItem) (models.CartItem, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCartItems(cartID int) ([]models.CartItem, error) {
	var cartItems []models.CartItem

	err := r.db.Where("cart_id=?", cartID).Preload("Cart").Preload("Cart.User").Find(&cartItems).Error

	return cartItems, err
}

func (r *repository) GetBook(bookID int) (models.Book, error) {
	var book models.Book

	err := r.db.Where("id=?", bookID).Preload("User").Find(&book).Error

	return book, err
}

func (r *repository) GetUser(userID int) (models.User, error) {
	var user models.User

	err := r.db.Where("id=?", userID).Find(&user).Error

	return user, err
}

func (r *repository) CreateCartItem(cartItem models.CartItem) (models.CartItem, error) {
	err := r.db.Create(&cartItem).Error

	return cartItem, err
}