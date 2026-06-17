package models

import "gorm.io/gorm"

type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model

	Name	 string `json:"name"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	Role	 string `json:"role"`
}