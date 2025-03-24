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
	exists, err := s.repo.CheckAdminExists()
	if err != nil {
		log.Fatalln("Failed to check if admin exists:", err)
	}
	
	if exists {
		log.Println("Admin already exists, skipping insert.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Failed to hash password:", err)
	}

	admin.Password = string(hashedPassword)

	err = s.repo.InsertAdmin(admin)
	if err != nil {
		log.Fatalln("Failed to insert admin:", err)
	}

	log.Println("Admin successfully inserted.")
}
