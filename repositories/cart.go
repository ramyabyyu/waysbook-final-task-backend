package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCartsByUserID(userID int) ([]models.Cart, error)
	AddCart(cart models.Cart) (models.Cart, error)
	GetCartByID(ID int) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCartsByUserID(userID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Where("user_id=?", userID).Preload("User").Preload("Book").Find(&carts).Error

	return carts, err
}

func (r *repository) AddCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) GetCartByID(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, "id=?", ID).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}