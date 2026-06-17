package config

import (
	"fmt"
	"jobconnect-api/models"
)

func Migration() {

	err := DB.AutoMigrate(
		&models.User{},
		&models.Job{},
	)

	if err != nil {
		fmt.Println("Migration gagal")
		return
	}

	fmt.Println("Migration berhasil")
}