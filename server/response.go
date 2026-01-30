package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, res *BaseResponse) {
	response := BaseResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    nil,
	}

	if res.Status != 0 {
		response.Status = res.Status
	}
	if res.Message != "" {
		response.Message = res.Message
	}
	if res.Data != nil {
		response.Data = res.Data
	}

	c.JSON(response.Status, response)
}
