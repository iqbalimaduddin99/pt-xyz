package usecases

import (
	"log"
	"pt-xyz/internal/entities"
	"pt-xyz/internal/repository"

	"golang.org/x/crypto/bcrypt"
)


type ServiceAdmin struct {
	repo repository.RepositoryAdmin
}

func NewServiceAdmin(repo repository.RepositoryAdmin) *ServiceAdmin {
	return &ServiceAdmin{repo: repo}
}

func (s *ServiceAdmin) AddAdmin(admin *entities.Admin) {
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
