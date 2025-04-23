package http

import (
	"net/http"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/usecases"
	"github.com/gin-gonic/gin"
)

type HandlerConsumer struct {
	service usecases.ServiceConsumer
}

func NewHandlerConsumer(service usecases.ServiceConsumer) *HandlerConsumer {
	return &HandlerConsumer{service: service}
}

func (h *HandlerConsumer) Register(c *gin.Context) {
	var consumer entities.ReqConsumer
	if err := c.ShouldBindJSON(&consumer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if len(consumer.KTP) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "KTP must not null"})
		return
	}

	if len(consumer.UserName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user name must not null"})
		return
	}

	if len(consumer.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must not null"})
		return
	}

	userName, err := h.service.RegisterConsumer(&consumer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userName": userName, "message": "User registered successfully"})
}




func (h *HandlerConsumer) Login(c *gin.Context) {
	var consumerReq entities.LoginRequest
	if err := c.ShouldBindJSON(&consumerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.service.Login(&consumerReq)

	if token == "Invalid username or password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err != nil  {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}  


	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}