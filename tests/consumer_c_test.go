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
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/register", h.Register)

	consumer := &entities.ReqConsumer{
		KTP:      "1234567890123456",
		UserName: "testuser",
		Password: "password123",
	}

	mockService.On("RegisterConsumer", consumer).Return(consumer.UserName, nil)

	body, _ := json.Marshal(consumer)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"userName":"testuser"`)
	mockService.AssertExpectations(t)
}

func TestRegisterHandler_InvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/register", h.Register)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(`invalid-json`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertNotCalled(t, "RegisterConsumer")
}

func TestRegisterHandler_KTPNull(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/register", h.Register)

	consumer := &entities.ReqConsumer{
		UserName: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(consumer)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "KTP must not null")
	mockService.AssertNotCalled(t, "RegisterConsumer")
}

func TestRegisterHandler_UserNameNull(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/register", h.Register)

	consumer := &entities.ReqConsumer{
		KTP:      "1234567890123456",
		Password: "password123",
	}

	body, _ := json.Marshal(consumer)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "user name must not null")
	mockService.AssertNotCalled(t, "RegisterConsumer")
}

func TestRegisterHandler_PasswordTooShort(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/register", h.Register)

	consumer := &entities.ReqConsumer{
		KTP:      "1234567890123456",
		UserName: "testuser",
		Password: "short",
	}

	body, _ := json.Marshal(consumer)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "password must not null")
	mockService.AssertNotCalled(t, "RegisterConsumer")
}

func TestRegisterHandler_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/register", h.Register)

	consumer := &entities.ReqConsumer{
		KTP:      "1234567890123456",
		UserName: "testuser",
		Password: "password123",
	}

	mockService.On("RegisterConsumer", consumer).Return("", errors.New("service error"))

	body, _ := json.Marshal(consumer)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), "service error")
	mockService.AssertExpectations(t)
}




func TestLoginHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/login", h.Login)

	consumerReq := &entities.LoginRequest{
		UserName: "testuser",
		Password: "password123",
	}

	mockService.On("Login", consumerReq).Return("valid-token", nil)

	body, _ := json.Marshal(consumerReq)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"message":"Login successful"`)
	assert.Contains(t, resp.Body.String(), `"token":"valid-token"`)
	mockService.AssertExpectations(t)
}

func TestLoginHandler_InvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/login", h.Login)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`invalid-json`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertNotCalled(t, "Login")
}

func TestLoginHandler_InvalidUsernameOrPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/login", h.Login)

	consumerReq := &entities.LoginRequest{
		UserName: "testuser",
		Password: "wrongpassword",
	}

	mockService.On("Login", consumerReq).Return("Invalid username or password", nil)

	body, _ := json.Marshal(consumerReq)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid username or password")
	mockService.AssertExpectations(t)
}

func TestLoginHandler_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockServiceConsumer)
	h := deliveryHttp.NewHandlerConsumer(mockService)

	router := gin.Default()
	router.POST("/login", h.Login)

	consumerReq := &entities.LoginRequest{
		UserName: "testuser",
		Password: "password123",
	}

	mockService.On("Login", consumerReq).Return("", errors.New("service error"))

	body, _ := json.Marshal(consumerReq)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "service error")
	mockService.AssertExpectations(t)
}