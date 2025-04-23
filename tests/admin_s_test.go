package tests

import (
	"errors"
	"testing"

	"pt-xyz/internal/entities"
	"pt-xyz/internal/usecases"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func TestAddAdmin_AdminAlreadyExists(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)

	mockAdminRepo.On("GetAdmin").Return(true, nil)

	svc := usecases.NewServiceAdmin(
		mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo,
	)

	admin := &entities.Admin{UserName: "admin", Password: "password"}
	svc.AddAdmin(admin)

	mockAdminRepo.AssertCalled(t, "GetAdmin")
	mockAdminRepo.AssertNotCalled(t, "CreateAdmin")
}

func TestAddAdmin_CheckAdminFails(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)


	mockAdminRepo.On("GetAdmin").Return(false, errors.New("db error"))
	mockAdminRepo.On("CreateAdmin", mock.Anything).Return(nil)

	svc := usecases.NewServiceAdmin(
		mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo,
	)

	admin := &entities.Admin{UserName: "admin", Password: "password"}
	svc.AddAdmin(admin)

	mockAdminRepo.AssertCalled(t, "GetAdmin")
	mockAdminRepo.AssertCalled(t, "CreateAdmin", mock.Anything)
}


func TestAddAdmin_CreateAdminFails(t *testing.T) {	
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)
	
	mockAdminRepo.On("GetAdmin").Return(false, nil)
	mockAdminRepo.On("CreateAdmin", mock.Anything).Return(errors.New("insert error"))

	svc := usecases.NewServiceAdmin(
		mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo,
	)

	admin := &entities.Admin{UserName: "admin", Password: "password"}
	svc.AddAdmin(admin)

	mockAdminRepo.AssertCalled(t, "CreateAdmin", mock.Anything)
}

func TestAddAdmin_Success(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)
	
	mockAdminRepo.On("GetAdmin").Return(false, nil)
	mockAdminRepo.On("CreateAdmin", mock.Anything).Return(nil)
	
	svc := usecases.NewServiceAdmin(
		mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo,
	)

	admin := &entities.Admin{UserName: "admin", Password: "password", FullName: "Admin XYZ"}
	svc.AddAdmin(admin)

	mockAdminRepo.AssertCalled(t, "GetAdmin")
	mockAdminRepo.AssertCalled(t, "CreateAdmin", mock.MatchedBy(func(a *entities.Admin) bool {
		return a.UserName == "admin" &&
			a.FullName == "Admin XYZ" &&
			a.Password != "password"
	}))
}


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
	mockDB := new(MockDB)
	service := usecases.NewServiceAdmin(mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo)

	id := uuid.New()

	mockAdminRepo.On("GetAdminByID", id).Return((*entities.Admin)(nil), errors.New("db error"))

	result, err := service.GetCreation(id)
	assert.Nil(t, result)
	assert.ErrorContains(t, err, "db error")
}


func TestAddLimitConsumer_Success(t *testing.T) {
	mockAdminRepo := new(MockRepoAdmin)
	mockProductRepo := new(MockRepoProduct)
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)

	idConsumer := uuid.New()

	loan := &entities.LoanLimit{}
	loan.ConsumerID=idConsumer
	
	consumer := &entities.Consumer{}
	consumer.ID=idConsumer

	mockRepoConsumer.On("GetConsumerById", idConsumer).Return(consumer, nil)
	mockRepoLoanLimit.On("GetLoanLimitByID", mockDB, idConsumer).Return((*entities.LoanLimit)(nil), nil)
	mockRepoLoanLimit.On("CreateLoanLimit", loan).Return(nil)

	svc := usecases.NewServiceAdmin(mockDB, mockAdminRepo, mockRepoConsumer, mockRepoLoanLimit, mockProductRepo)

	result, err := svc.AddLimitConsumer(loan)

	assert.NoError(t, err)
	assert.Equal(t, loan.ConsumerID.String(), result)

	mockRepoConsumer.AssertExpectations(t)
	mockRepoLoanLimit.AssertExpectations(t)
}

func TestAddLimitConsumer_ConsumerNotFound(t *testing.T) {
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)

	idConsumer := uuid.New()
	idLoan := uuid.New()
	loan := &entities.LoanLimit{}
	loan.ID = idLoan
	loan.ConsumerID = idConsumer

	mockRepoConsumer.On("GetConsumerById", idConsumer).Return(&entities.Consumer{}, errors.New("consumer not found"))

	svc := usecases.NewServiceAdmin(mockDB, nil, mockRepoConsumer, mockRepoLoanLimit, nil)

	result, err := svc.AddLimitConsumer(loan)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, "consumer not found", err.Error())

	mockRepoConsumer.AssertExpectations(t)
}


func TestAddLimitConsumer_ConsumerAlreadyHasLimit(t *testing.T) {
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)

	idConsumer := uuid.New()
	loan := &entities.LoanLimit{
		ConsumerID: idConsumer,
		LimitLoan:      10000,
	}

	consumer := &entities.Consumer{}
	consumer.ID=idConsumer

	mockRepoConsumer.On("GetConsumerById", idConsumer).Return(consumer, nil)
	mockRepoLoanLimit.On("GetLoanLimitByID", mockDB, idConsumer).Return(&entities.LoanLimit{}, nil)

	svc := usecases.NewServiceAdmin(mockDB, nil, mockRepoConsumer, mockRepoLoanLimit, nil)

	result, err := svc.AddLimitConsumer(loan)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, "Consumer have limit", err.Error())

	mockRepoConsumer.AssertExpectations(t)
	mockRepoLoanLimit.AssertExpectations(t)
}

func TestAddLimitConsumer_CreateLoanLimitFails(t *testing.T) {
	mockRepoConsumer := new(MockRepoConsumer)
	mockRepoLoanLimit := new(MockRepoLoanLimit)
	mockDB := new(MockDB)

	idConsumer := uuid.New()
	loan := &entities.LoanLimit{
		ConsumerID: idConsumer,
		LimitLoan:      10000,
	}

	consumer := &entities.Consumer{}
	consumer.ID=idConsumer

	mockRepoConsumer.On("GetConsumerById", idConsumer).Return(consumer, nil)
	mockRepoLoanLimit.On("GetLoanLimitByID", mockDB, idConsumer).Return((*entities.LoanLimit)(nil), nil)
	mockRepoLoanLimit.On("CreateLoanLimit", loan).Return(errors.New("failed to insert loan limit"))

	svc := usecases.NewServiceAdmin(mockDB, nil, mockRepoConsumer, mockRepoLoanLimit, nil)

	result, err := svc.AddLimitConsumer(loan)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, "failed to insert loan limit: failed to insert loan limit", err.Error())

	mockRepoConsumer.AssertExpectations(t)
	mockRepoLoanLimit.AssertExpectations(t)
}
