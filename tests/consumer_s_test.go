package tests

import (
	"errors"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/usecases"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterConsumer_Success(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	service := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	consumer := &entities.ReqConsumer{
		UserName: "testuser",
		Password: "password123",
		KTP:      "1234567890123456",
	}

	mockRepo.On("GetConsumerByUserName", consumer.UserName).Return((*entities.Consumer)(nil), nil)
	mockRepo.On("GetConsumerByKTP", consumer.KTP).Return((*entities.Consumer)(nil), nil)
	mockRepo.On("CreateConsumer", consumer).Return(nil)

	mockRepo.On("CreateConsumer", mock.Anything).Return(nil)

	userName, err := service.RegisterConsumer(consumer)

	assert.NoError(t, err)
	assert.Equal(t, "testuser", userName)
	mockRepo.AssertExpectations(t)
}

func TestRegisterConsumer_UserExists(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	service := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	consumer := &entities.ReqConsumer{
		UserName: "testuser",
		Password: "password123",
		KTP:      "1234567890123456",
	}

	mockRepo.On("GetConsumerByUserName", consumer.UserName).Return(&entities.Consumer{}, nil)
	mockRepo.On("GetConsumerByKTP", consumer.KTP).Return((*entities.Consumer)(nil), nil)

	userName, err := service.RegisterConsumer(consumer)

	assert.Error(t, err)
	assert.Equal(t, "", userName)
	assert.Contains(t, err.Error(), "Consumer already exists")
	mockRepo.AssertExpectations(t)
}

func TestRegisterConsumer_KTPExists(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	service := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	consumer := &entities.ReqConsumer{
		UserName: "testuser",
		Password: "password123",
		KTP:      "1234567890123456",
	}

	mockRepo.On("GetConsumerByUserName", consumer.UserName).Return((*entities.Consumer)(nil), nil)
	mockRepo.On("GetConsumerByKTP", consumer.KTP).Return(&entities.Consumer{}, nil)

	userName, err := service.RegisterConsumer(consumer)

	assert.Error(t, err)
	assert.Equal(t, "", userName)
	assert.Contains(t, err.Error(), "Consumer already exists")
	mockRepo.AssertExpectations(t)
}

func TestRegisterConsumer_FailToInsertConsumer(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	service := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	consumer := &entities.ReqConsumer{
		UserName: "testuser",
		Password: "password123",
		KTP:      "1234567890123456",
	}

	mockRepo.On("GetConsumerByUserName", consumer.UserName).Return((*entities.Consumer)(nil), nil)
	mockRepo.On("GetConsumerByKTP", consumer.KTP).Return((*entities.Consumer)(nil), nil)

	mockRepo.On("CreateConsumer", consumer).Return(errors.New("Failed to insert consumer"))

	userName, err := service.RegisterConsumer(consumer)

	assert.Error(t, err)
	assert.Equal(t, "", userName)
	assert.Contains(t, err.Error(), "Failed to insert consumer")
	mockRepo.AssertExpectations(t)
}









func hashPassword(pw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash)
}

func TestLoginConsumer_Success(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	svc := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	hashed := hashPassword("password123")
	uid := uuid.New()

	consumer := &entities.Consumer{
		UserName: "testuser",
		Password: hashed,
	}
	consumer.ID=uid

	mockRepo.On("GetConsumerByUserName", "testuser").Return(consumer, nil)

	req := &entities.LoginRequest{
		UserName: "testuser",
		Password: "password123",
		IsAdmin:  false,
	}

	token, err := svc.Login(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLoginConsumer_InvalidPassword(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	svc := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	hashed := hashPassword("correctpassword")
	
	consumer := &entities.Consumer{
		UserName: "testuser",
		Password: hashed,
	}
	consumer.ID=uuid.New()
	mockRepo.On("GetConsumerByUserName", "testuser").Return(consumer, nil)

	req := &entities.LoginRequest{
		UserName: "testuser",
		Password: "wrongpassword",
		IsAdmin:  false,
	}

	_, err := svc.Login(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid username or password")
	mockRepo.AssertExpectations(t)
}

func TestLoginConsumer_UserNotFound(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	svc := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	mockRepo.On("GetConsumerByUserName", "notfound").Return((*entities.Consumer)(nil), nil)

	req := &entities.LoginRequest{
		UserName: "notfound",
		Password: "password123",
		IsAdmin:  false,
	}

	_, err := svc.Login(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid username or password")
	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin_Success(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	svc := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	hashed := hashPassword("adminpass")
	uid := uuid.New()

	admin := &entities.Admin{
		UserName: "adminuser",
		Password: hashed,
	}
	admin.ID=uid

	mockRepoAdmin.On("GetAdminByUserName", "adminuser").Return(admin, nil)

	req := &entities.LoginRequest{
		UserName: "adminuser",
		Password: "adminpass",
		IsAdmin:  true,
	}

	token, err := svc.Login(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepoAdmin.AssertExpectations(t)
}

func TestLoginAdmin_InvalidPassword(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	svc := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	hashed := hashPassword("correctadminpass")

	
	admin := &entities.Admin{
		UserName: "adminuser",
		Password: hashed,
	}
	admin.ID=uuid.New()
	mockRepoAdmin.On("GetAdminByUserName", "adminuser").Return(admin, nil)

	req := &entities.LoginRequest{
		UserName: "adminuser",
		Password: "wrongadminpass",
		IsAdmin:  true,
	}

	_, err := svc.Login(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid username or password")
	mockRepoAdmin.AssertExpectations(t)
}

func TestLoginAdmin_UserNotFound(t *testing.T) {
	mockRepo := new(MockRepoConsumer)
	mockRepoAdmin := new(MockRepoAdmin)

	svc := usecases.NewServiceConsumer(mockRepo, mockRepoAdmin)

	mockRepoAdmin.On("GetAdminByUserName", "notfound").Return((*entities.Admin)(nil), nil)

	req := &entities.LoginRequest{
		UserName: "notfound",
		Password: "adminpass",
		IsAdmin:  true,
	}

	_, err := svc.Login(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid username or password")
	mockRepoAdmin.AssertExpectations(t)
}