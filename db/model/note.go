package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Judul     string `form:"judul" json:"judul" binding:"required" gorm:"not null"`
	Deskripsi string `form:"deskripsi" json:"deskripsi" binding:"required"`
}
