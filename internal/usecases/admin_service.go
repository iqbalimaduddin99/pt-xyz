package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"pt-xyz/configs/database"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServiceAdmin interface {
	AddAdmin(admin *entities.Admin)
	AddLimitConsumer(loan *entities.LoanLimit) (string, error)
	GetCreation(id uuid.UUID) (*entities.MasterProductPtXyz, error)
}

type ServiceAdminImpl struct {
	db database.Database
	repo repository.RepositoryAdmin
	repoProduct repository.RepositoryMasterProductXYZ
	repoConsumer repository.RepositoryConsumer
	repoLoanLimit repository.RepositoryLoanLimit
}

func NewServiceAdmin(db database.Database, repo repository.RepositoryAdmin, repoConsumer repository.RepositoryConsumer, repoLoanLimit repository.RepositoryLoanLimit, repoProduct repository.RepositoryMasterProductXYZ) *ServiceAdminImpl {
	return &ServiceAdminImpl{ db:db, repo: repo, repoProduct: repoProduct, repoConsumer: repoConsumer, repoLoanLimit: repoLoanLimit}
}

func (s *ServiceAdminImpl) AddAdmin(admin *entities.Admin) {
	exists, err := s.repo.GetAdmin()
	if err != nil {
		log.Println("Failed to check if admin exists")
	}
	
	if exists {
		log.Println("Admin already exists, skipping insert.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err)
	}

	admin.Password = string(hashedPassword)

	err = s.repo.CreateAdmin(admin)
	if err != nil {
		log.Println("Failed to insert admin:", err)
	}

	log.Println("Admin successfully inserted.")
}



func (s *ServiceAdminImpl) GetCreation(id uuid.UUID) (*entities.MasterProductPtXyz, error) {
	admin, err := s.repo.GetAdminByID(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if admin == nil {
		return nil, fmt.Errorf("Invalid username or password")
	}

	adminCreation, err := s.repoProduct.GetMasterProductByCreator(id)
	
	return adminCreation, nil
}


func (s *ServiceAdminImpl) AddLimitConsumer(loan *entities.LoanLimit) (string, error) {
	_, err := s.repoConsumer.GetConsumerById(loan.ConsumerID)
	if err != nil {
			return "", err
    }

	fmt.Print(s.db)
	existsLimitLoan, err := s.repoLoanLimit.GetLoanLimitByID(s.db, loan.ConsumerID)
	if err != nil && err != sql.ErrNoRows  {
			return "", err
    }

	fmt.Println("err", err)
	fmt.Println("existsLimitLoan", existsLimitLoan)
	if existsLimitLoan != nil {
		return "", errors.New("Consumer have limit")
	}

	err = s.repoLoanLimit.CreateLoanLimit(loan)
	if err != nil {
		return "", fmt.Errorf("failed to insert loan limit: %w", err)
	}
	
	return loan.ConsumerID.String(), nil
}
