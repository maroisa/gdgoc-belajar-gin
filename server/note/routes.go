package note

import (
	"belajar-gin/db/model"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (nh *NoteHandler) Routes() {
	note := nh.r.Group("/note")
	note.GET("/", nh.getNotes)
	note.POST("/", nh.postNotes)
	note.PUT("/:id", nh.putNotes)
	note.DELETE("/:id", nh.deleteNotes)
}

func (nh *NoteHandler) getNotes(c *gin.Context) {
	ctx := context.Background()

	notes, err := gorm.G[model.Note](nh.db).Find(ctx)
	if err != nil {
		log.Println("Failed to get notes: ", err.Error())
	}

	log.Println(notes)
}

func (nh *NoteHandler) postNotes(c *gin.Context) {

}

func (nh *NoteHandler) putNotes(c *gin.Context) {

}

func (nh *NoteHandler) deleteNotes(c *gin.Context) {

}
