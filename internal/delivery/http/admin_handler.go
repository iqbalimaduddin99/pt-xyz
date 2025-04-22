package http

import (
	"net/http"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/usecases"
	"pt-xyz/pkg"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerAdmin struct {
	service usecases.ServiceAdmin
}

func NewHandlerAdmin(service usecases.ServiceAdmin) *HandlerAdmin {
	return &HandlerAdmin{service: service}
}

func (h *HandlerAdmin) GetCreation(c *gin.Context) {
	id := c.Param("id")
	idParse, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
        return
    }

	creation, err := h.service.GetCreation(idParse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pkg.Success(c, creation, nil, "success")
}



func (h *HandlerAdmin) AddLimitConsumer(c *gin.Context) {
	var loanLimit entities.LoanLimit
	if err := c.ShouldBindJSON(&loanLimit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	consumer, err := h.service.AddLimitConsumer(&loanLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pkg.Success(c, consumer, nil, "success")
}
