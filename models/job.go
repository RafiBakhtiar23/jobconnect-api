package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model

	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Company   string `json:"company" binding:"required"`
	Lokasi    string `json:"lokasi" binding:"required"`
	Gaji      int    `json:"gaji" binding:"required"`
}