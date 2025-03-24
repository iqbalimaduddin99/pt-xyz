package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}, meta interface{}, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Data:    data,
		Meta:    meta,
		Status:  "success",
		Code:    http.StatusOK,
		Message: message,
	})
}

func Fail(c *gin.Context, code int, message string, errorDetails interface{}) {
	c.JSON(code, APIResponse{
		Data:    nil,
		Meta:    nil,
		Status:  "fail",
		Code:    code,
		Message: message,
		Error:   errorDetails,
	})
}

func Error(c *gin.Context, message string, errorDetails interface{}) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Data:    nil,
		Meta:    nil,
		Status:  "error",
		Code:    http.StatusInternalServerError,
		Message: message,
		Error:   errorDetails,
	})
}
