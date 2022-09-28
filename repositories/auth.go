package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	GetUser(ID int) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetOldUserPassword(email string) (string, error)
	CheckEmailExist(email string) error
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User

	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User

	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetOldUserPassword(email string) (string, error) {
	var user models.User

	err := r.db.Raw("SELECT password FROM users WHERE email=?", email).Error
	return user.Password, err
}

func (r *repository) CheckEmailExist(email string) error {
	var user models.User

	err := r.db.First(&user, "email=?", email).Error

	return err
}