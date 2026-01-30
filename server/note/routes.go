package note

import "github.com/gin-gonic/gin"

func (nh *NoteHandler) Routes() {
	note := nh.r.Group("/note")
	note.GET("/", getNotes)
	note.POST("/", postNotes)
	note.PUT("/:id", putNotes)
	note.DELETE("/:id", deleteNotes)
}

func getNotes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func postNotes(c *gin.Context) {

}

func putNotes(c *gin.Context) {

}

func deleteNotes(c *gin.Context) {

}
