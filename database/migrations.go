package database

import (
	"fmt"
	"waysbook/models"
	psql "waysbook/pkg/dbConnection"
)

func RunMigration() {
	err := psql.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}