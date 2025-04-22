package tests

import (
	"database/sql"
	"pt-xyz/configs/database"
	"pt-xyz/internal/entities"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type MockRepoAdmin struct{ mock.Mock }
type MockRepoConsumer struct{ mock.Mock }
type MockRepoLoanLimit struct{ mock.Mock }
type MockRepoProduct struct{ mock.Mock }
type MockDB struct { mock.Mock }

func (m *MockDB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	argsList := m.Called(query, args)
	return argsList.Get(0).(*sqlx.Rows), argsList.Error(1)
}

func (m *MockDB) Get(dest interface{}, query string, args ...interface{}) error {
	argsList := m.Called(dest, query, args)
	return argsList.Error(0)
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	argsList := m.Called(query, args)
	return argsList.Get(0).(sql.Result), argsList.Error(1)
}

func (m *MockDB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	argsList := m.Called(query, arg)
	return argsList.Get(0).(sql.Result), argsList.Error(1)
}

func (m *MockDB) MustBegin() *sqlx.Tx {
	argsList := m.Called()
	return argsList.Get(0).(*sqlx.Tx)
}

func (m *MockDB) Ping() error {
	argsList := m.Called()
	return argsList.Error(0)
}

func (m *MockDB) SetMaxOpenConns(n int) {
	m.Called(n)
}

func (m *MockDB) SetMaxIdleConns(n int) {
	m.Called(n)
}

func (m *MockDB) SetConnMaxLifetime(d time.Duration) {
	m.Called(d)
}

func (m *MockDB) Close() error {
	argsList := m.Called()
	return argsList.Error(0)
}

// MockRepoAdmin
func (m *MockRepoAdmin) GetAdmin() (bool, error)                            { return true, nil }
func (m *MockRepoAdmin) CreateAdmin(*entities.Admin) error                  { return nil }
func (m *MockRepoAdmin) GetAdminByUserName(string) (*entities.Admin, error) { return nil, nil }

func (m *MockRepoAdmin) GetAdminByID(id uuid.UUID) (*entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Admin), args.Error(1)
}


// MockRepoConsumer
func (m *MockRepoConsumer) GetConsumerById(id uuid.UUID) (*entities.Consumer, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Consumer), args.Error(1)
}

func (m *MockRepoConsumer) GetConsumerByUserName(userName string) (*entities.Consumer, error) {
	args := m.Called(userName)
	return args.Get(0).(*entities.Consumer), args.Error(1)
}

func (m *MockRepoConsumer) GetConsumerByKTP(KTP string) (*entities.Consumer, error) {
	args := m.Called(KTP)
	return args.Get(0).(*entities.Consumer), args.Error(1)
}

func (m *MockRepoConsumer) CreateConsumer(consumer *entities.ReqConsumer) error {
	args := m.Called(consumer)
	return args.Error(0)
}

//MockRepoLoanLimit

func (m *MockRepoLoanLimit) GetLoanLimitByID(tx database.Database, consumerID uuid.UUID) (*entities.LoanLimit, error) {
	args := m.Called(tx, consumerID)
	return args.Get(0).(*entities.LoanLimit), args.Error(1)
}

func (m *MockRepoLoanLimit) CreateLoanLimit(limit *entities.LoanLimit) error {
	args := m.Called(limit)
	return args.Error(0)
}

// MockRepoProduct
func (m *MockRepoProduct) GetMasterProductForTransactionById(tx *sqlx.DB, id uuid.UUID) (*entities.MasterProductPtXyz, error) {
	panic("unimplemented")
}
func (m *MockRepoProduct) GetMasterProductByCreator(id uuid.UUID) (*entities.MasterProductPtXyz, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.MasterProductPtXyz), args.Error(1)
}