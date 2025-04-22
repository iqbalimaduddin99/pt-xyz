package http

import (
	"net/http"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/usecases"
	"pt-xyz/pkg"

	"github.com/gin-gonic/gin"
)

type HandlerTransaction struct {
	service *usecases.ServiceTransaction
}

func NewHandlerTransaction(service *usecases.ServiceTransaction) *HandlerTransaction {
	return &HandlerTransaction{service: service}
}
func (h *HandlerTransaction) CreateTransaction(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		pkg.Fail(c, http.StatusUnauthorized, "Authentication required", nil)
		c.Abort()
		return
	}

	typedClaims, _ := claims.(*pkg.Claims)

	var transaction entities.TransactionTableReq
	if err := c.ShouldBindJSON(&transaction); err != nil {
		pkg.Fail(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}
	
	message, err := h.service.CreateTransaction(&transaction, typedClaims)
	if err != nil {
		pkg.Error(c, "Error processing transaction", err.Error())
		return
	}

	pkg.Success(c, nil, nil, message)
}


