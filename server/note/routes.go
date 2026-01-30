package note

import (
	"belajar-gin/db/model"
	"belajar-gin/server"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (nh *NoteHandler) Routes() {
	note := nh.r.Group("/note")
	note.GET("", nh.getNotes)
	note.POST("", nh.postNotes)
	note.PUT("/:id", nh.putNote)
	note.GET("/:id", nh.getNote)
	note.DELETE("/:id", nh.deleteNotes)
}

func (nh *NoteHandler) getNotes(c *gin.Context) {
	notes, err := gorm.G[model.Note](nh.db).Find(context.Background())
	if err != nil {
		server.Response(c, &server.BaseResponse{
			Message: err.Error(),
		})
		return
	}
	server.Response(c, &server.BaseResponse{
		Data: notes,
	})
}

func (nh *NoteHandler) getNote(c *gin.Context) {
	user, err := gorm.G[model.Note](nh.db).Where("id = ?", c.Param("id")).First(context.Background())
	if err != nil {
		server.Response(c, &server.BaseResponse{
			Status:  404,
			Message: "Not Found",
		})
		return
	}

	server.Response(c, &server.BaseResponse{
		Data: user,
	})
}

func (nh *NoteHandler) postNotes(c *gin.Context) {
	var note model.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		server.Response(c, &server.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
		return
	}

	if note.Judul == "" || note.Deskripsi == "" {
		server.Response(c, &server.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Field cannot be empty",
		})
		return
	}

	err := gorm.G[model.Note](nh.db).Create(context.Background(), &note)
	if err != nil {
		server.Response(c, &server.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to insert note: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, note)
}

func (nh *NoteHandler) putNote(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	var note model.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		server.Response(c, &server.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
		return
	}

	if note.Judul != "" {
		_, err := gorm.G[model.Note](nh.db).Where("id = ?", id).Update(ctx, "judul", note.Judul)
		if err != nil {
			server.Response(c, &server.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Bad Request",
			})
			return
		}
	}

	if note.Deskripsi != "" {
		_, err := gorm.G[model.Note](nh.db).Where("id = ?", id).Update(ctx, "deskripsi", note.Deskripsi)
		if err != nil {
			server.Response(c, &server.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Bad Request",
			})
			return
		}
	}

	server.Response(c, &server.BaseResponse{})
}

func (nh *NoteHandler) deleteNotes(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	rowsAffected, err := gorm.G[model.Note](nh.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		server.Response(c, &server.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
		return
	}

	if rowsAffected == 0 {
		server.Response(c, &server.BaseResponse{
			Status:  404,
			Message: "Not Found",
		})
		return
	}

	server.Response(c, &server.BaseResponse{})
}
