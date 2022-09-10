package database

import (
	"fmt"

	"github.com/ZaphCode/go-fiber-ps/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.Product{}, &models.Review{})

	if err != nil {
		fmt.Println(">>> Migration Error: " + err.Error())
		panic("Migration Fail")
	}

	fmt.Println("Models migrated")
}
