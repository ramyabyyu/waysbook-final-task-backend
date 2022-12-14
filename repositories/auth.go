package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	Login(email string) (models.User, error)
	CheckEmailExist(email string) (error)
	BecomeSeller(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User

	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CheckEmailExist(email string) (error) {
	var user models.User

	err := r.db.First(&user, "email=?", email).Error

	return err
}

func (r *repository) BecomeSeller(ID int) (models.User, error) {
	var user models.User
	r.db.First(&user, "id=?", ID)

	// Change user IsSeller to true
	user.IsSeller = true
	err := r.db.Save(&user).Error

	return user, err
}