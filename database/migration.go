package database

import (
	"backend-api/models"
	"backend-api/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Singer{},
		&models.Music{},
		&models.Transaction{},
		// &models.Subscription{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration failed.!")
	}

	fmt.Println("Migration successfully")
}
