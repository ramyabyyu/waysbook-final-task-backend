package models

import "time"

type User struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	FullName string `json:"full_name" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" gorm:"type: varchar(255)"`
	IsSeller  bool   `json:"is_seller"`
	Gender string `json:"gender" gorm:"type: varchar(255)"`
	Phone string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: text"`
	Photo string `json:"photo"`
	IsPhotoChange bool `json:"is_photo_change"`
	Books []Book `json:"books"`
	Carts []Cart `json:"carts"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
