package tests

import (
	"database/sql"
	"errors"
	"testing"

	"pt-xyz/internal/entities"
	"pt-xyz/internal/usecases"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)



func TestGetCreation_Success(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)
	service := usecases.NewServiceAdmin(mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo)

	id := uuid.New()
	admin := &entities.Admin{}
	admin.ID=id
	product := &entities.MasterProductPtXyz{}
	uid, _ := uuid.Parse("a32f0f36-3c88-4f3b-91d0-b2e8f3c0c4aa")
	product.ID=uid

	mockAdminRepo.On("GetAdminByID", id).Return(admin, nil)
	mockProductRepo.On("GetMasterProductByCreator", id).Return(product, nil)

	result, err := service.GetCreation(id)
	assert.NoError(t, err)
	assert.Equal(t, "a32f0f36-3c88-4f3b-91d0-b2e8f3c0c4aa", result.ID.String())
}

func TestGetCreation_AdminNotFound(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)
	service := usecases.NewServiceAdmin(mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo)

	id := uuid.New()

	mockAdminRepo.On("GetAdminByID", id).Return((*entities.Admin)(nil), nil)

	result, err := service.GetCreation(id)
	assert.Nil(t, result)
	assert.ErrorContains(t, err, "Invalid username or password")
}

func TestGetCreation_DBError(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	service := usecases.NewServiceAdmin(sqlx.NewDb(&sql.DB{},""), mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo)

	id := uuid.New()

	mockAdminRepo.On("GetAdminByID", id).Return((*entities.Admin)(nil), errors.New("db error"))

	result, err := service.GetCreation(id)
	assert.Nil(t, result)
	assert.ErrorContains(t, err, "db error")
}
