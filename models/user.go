package models

import "time"

type User struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" gorm:"type: varchar(255)"`
	IsSeller  bool   `json:"is_seller"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}