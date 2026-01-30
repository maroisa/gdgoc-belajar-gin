package note

import (
	"belajar-gin/db/model"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (nh *NoteHandler) Routes() {
	note := nh.r.Group("/note")
	note.GET("", nh.getNotes)
	note.POST("", nh.postNotes)
	note.PUT("/:id", nh.putNotes)
	note.DELETE("/:id", nh.deleteNotes)
}

func (nh *NoteHandler) getNotes(c *gin.Context) {
	notes, err := gorm.G[model.Note](nh.db).Find(context.Background())
	if err != nil {
		log.Println("Failed")
		return
	}
	c.JSON(http.StatusOK, notes)
}

func (nh *NoteHandler) postNotes(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if note.Judul == "" || note.Deskripsi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := gorm.G[model.Note](nh.db).Create(context.Background(), &note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to insert note: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, note)
}

func (nh *NoteHandler) putNotes(c *gin.Context) {

}

func (nh *NoteHandler) deleteNotes(c *gin.Context) {

}
