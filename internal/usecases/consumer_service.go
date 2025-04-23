package usecases

import (
	"database/sql"
	"fmt"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/repository"
	"pt-xyz/pkg"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type ServiceConsumer interface {
	RegisterConsumer(consumer *entities.ReqConsumer) (string, error)
	Login(loginRequest *entities.LoginRequest) (string, error)
}

type ServiceConsumerImpl struct {
	repo repository.RepositoryConsumer
	repoAdmin repository.RepositoryAdmin
	jwtGen    func(userID uuid.UUID, userName string, isAdmin bool) (string, error)
}

func NewServiceConsumer(repo repository.RepositoryConsumer, repoAdmin repository.RepositoryAdmin) *ServiceConsumerImpl {
	return &ServiceConsumerImpl{repo: repo, repoAdmin: repoAdmin, jwtGen: pkg.GenerateJWT,}
}

func (s *ServiceConsumerImpl) RegisterConsumer(consumer *entities.ReqConsumer) (string, error) {
	exists, err := s.repo.GetConsumerByUserName(consumer.UserName)
	if err != nil && err != sql.ErrNoRows {
			return "", err
    }

	existsKTP, err := s.repo.GetConsumerByKTP(consumer.KTP)
	if err != nil && err != sql.ErrNoRows {
			return "", err
    }

	fmt.Println(exists)
	
	if exists != nil || existsKTP != nil {
		return "", fmt.Errorf("Consumer already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(consumer.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password:", err)
	}

	consumer.Password = string(hashedPassword)

	err = s.repo.CreateConsumer(consumer)
	if err != nil {
		return "", fmt.Errorf("Failed to insert consumer:", err)
	}

	return consumer.UserName, nil
}




func (s *ServiceConsumerImpl) Login(loginRequest *entities.LoginRequest) (string, error) {
	var token string
	var err error

	if loginRequest.IsAdmin {
		admin, err := s.repoAdmin.GetAdminByUserName(loginRequest.UserName)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}

		if admin == nil {
			return "", fmt.Errorf("Invalid username or password")
		}
		
		token, err = ComparePassAndGenerateJWT(admin.Password, loginRequest.Password, true, admin.ID, admin.UserName, s.jwtGen)
		if err != nil {
			return "", fmt.Errorf("Invalid username or password")
		}
	} else {
		consumer, err := s.repo.GetConsumerByUserName(loginRequest.UserName)
		if err != nil || consumer == nil {
			return "", fmt.Errorf("Invalid username or password")
		}
		
		token, err = ComparePassAndGenerateJWT(consumer.Password, loginRequest.Password, false, consumer.ID, consumer.UserName, s.jwtGen)
		if err != nil {
			return "", fmt.Errorf("Invalid username or password")
		}
	}

	if err != nil {
		return "", err
	}
	return token, nil
}


func ComparePassAndGenerateJWT(passwordEncrypted string, reqPassword string, isAdmin bool, userID uuid.UUID, userName string, genJWT func(uuid.UUID, string, bool) (string, error)) (string, error) {

	err := bcrypt.CompareHashAndPassword([]byte(passwordEncrypted), []byte(reqPassword))
	if err != nil {
		return "", fmt.Errorf("Invalid username or password")
	}

	
	token, err := genJWT(userID, userName, isAdmin)
	if err != nil {
		return "", fmt.Errorf("Error generating token: %v", err)
	}
	return token, nil
}
	
