package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Judul string
	Deskripsi string
}
