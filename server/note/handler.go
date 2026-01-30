package note

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NoteHandler struct {
	r  *gin.Engine
	db *gorm.DB
}

func NewHandler(r *gin.Engine, db *gorm.DB) *NoteHandler {
	nh := &NoteHandler{
		r:  r,
		db: db,
	}

	nh.Routes()

	return nh
}
