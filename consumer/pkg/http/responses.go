package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponseData struct {
	Status string      `json:"status"` // "success"
	Data   interface{} `json:"data"`
}

type ErrorResponseData struct {
	Status string `json:"status"` // "error"
	Error  string `json:"error"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponseData{
		Status: "success",
		Data:   data,
	})
}

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ErrorResponseData{
		Status: "error",
		Error:  err.Error(),
	})
}
