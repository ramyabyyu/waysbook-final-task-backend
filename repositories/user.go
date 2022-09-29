package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	GetUserByID(ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUserByID(ID int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, "id=?", ID).Error

	return user, err
}