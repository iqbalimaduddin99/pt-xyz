package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	deliveryHttp "pt-xyz/internal/delivery/http"
	"pt-xyz/internal/entities"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServiceAdmin struct{ 
	mock.Mock 
}

func (m *MockServiceAdmin) GetCreation(id uuid.UUID) (*entities.MasterProductPtXyz, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.MasterProductPtXyz), args.Error(1)
}

func (m *MockServiceAdmin) AddAdmin(admin *entities.Admin) {
	panic("unimplemented")
}


func (m *MockServiceAdmin) AddLimitConsumer(loan *entities.LoanLimit) (string, error) {
	panic("unimplemented")
}

func TestGetCreationHandler_Success(t *testing.T) {
	mockService := new(MockServiceAdmin)
	handler := deliveryHttp.NewHandlerAdmin(mockService)

	router := gin.Default()
	router.GET("/get-creation/:id", handler.GetCreation)

	id := uuid.New()
	mockProduct := &entities.MasterProductPtXyz{}
	uid, _ := uuid.Parse("a32f0f36-3c88-4f3b-91d0-b2e8f3c0c4aa")
	mockProduct.ID=uid

	mockService.On("GetCreation", id).Return(mockProduct, nil)

	req := httptest.NewRequest(http.MethodGet, "/get-creation/"+id.String(), nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), `"status":"success"`)
}

func TestGetCreationHandler_InvalidUUID(t *testing.T) {
	mockService := new(MockServiceAdmin)
	handler := deliveryHttp.NewHandlerAdmin(mockService)

	router := gin.Default()
	router.GET("/get-creation/:id", handler.GetCreation)

	req := httptest.NewRequest(http.MethodGet, "/get-creation/invalid-uuid", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid UUID")
}

func TestGetCreationHandler_ServiceError(t *testing.T) {
	mockService := new(MockServiceAdmin)
	handler := deliveryHttp.NewHandlerAdmin(mockService)

	router := gin.Default()
	router.GET("/get-creation/:id", handler.GetCreation)

	id := uuid.New()
	mockService.On("GetCreation", id).Return(nil, errors.New("mock error"))

	req := httptest.NewRequest(http.MethodGet, "/get-creation/"+id.String(), nil)
	resp := httptest.NewRecorder()
	t.Log("Response body:", resp.Body.String())

	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
	assert.Contains(t, resp.Body.String(), "error")
}
