package tests

import (
	"pt-xyz/internal/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)


type MockServiceAdmin struct{ mock.Mock }
type MockServiceConsumer struct{ mock.Mock }


func (m *MockServiceAdmin) GetCreation(id uuid.UUID) (*entities.MasterProductPtXyz, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.MasterProductPtXyz), args.Error(1)
}

func (m *MockServiceAdmin) AddAdmin(admin *entities.Admin) {
	m.Called(admin)
}

func (m *MockServiceAdmin) AddLimitConsumer(loan *entities.LoanLimit) (string, error) {
	args := m.Called(loan)
	return args.String(0), args.Error(1)
}


func (m *MockServiceConsumer) RegisterConsumer(consumer *entities.ReqConsumer) (string, error) {
	args := m.Called(consumer)
	return args.String(0), args.Error(1)
}

func (m *MockServiceConsumer) Login(loginRequest *entities.LoginRequest) (string, error) {
	args := m.Called(loginRequest)
	return args.String(0), args.Error(1)
}
