package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	deliveryHttp "pt-xyz/internal/delivery/http"
	"pt-xyz/internal/entities"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


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




func TestAddLimitConsumerHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceAdmin)
	h := deliveryHttp.NewHandlerAdmin(mockService)

	router := gin.Default()
	router.POST("/add-limit/consumer", h.AddLimitConsumer)

	consumerID := uuid.New()
	loan := &entities.LoanLimit{}
	loan.ConsumerID=consumerID
	loan.ID=uuid.New()
	loan.LimitLoan=50000

	mockService.On("AddLimitConsumer", loan).Return(consumerID.String(), nil)

	body, _ := json.Marshal(loan)
	req, _ := http.NewRequest(http.MethodPost, "/add-limit/consumer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestAddLimitConsumerHandler_InvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceAdmin)
	h := deliveryHttp.NewHandlerAdmin(mockService)

	router := gin.Default()
	router.POST("/add-limit/consumer", h.AddLimitConsumer)

	req, _ := http.NewRequest(http.MethodPost, "/add-limit/consumer", bytes.NewBufferString(`invalid-json`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertNotCalled(t, "AddLimitConsumer")
}

func TestAddLimitConsumerHandler_FailedToAdd(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceAdmin)
	h := deliveryHttp.NewHandlerAdmin(mockService)

	router := gin.Default()
	router.POST("/add-limit/consumer", h.AddLimitConsumer)

	consumerID := uuid.New()
	loan := &entities.LoanLimit{}
	loan.ConsumerID=consumerID
	loan.ID=uuid.New()
	loan.LimitLoan=50000

	mockService.On("AddLimitConsumer", loan).Return("", errors.New("something went wrong"))

	body, _ := json.Marshal(loan)
	req, _ := http.NewRequest(http.MethodPost, "/add-limit/consumer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}