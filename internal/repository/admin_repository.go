package repository

import (
	"pt-xyz/configs/database"
	"pt-xyz/internal/entities"

	"github.com/google/uuid"
)

type RepositoryAdmin interface {
	GetAdmin() (bool, error)
	CreateAdmin(admin *entities.Admin) error
	GetAdminByUserName(userName string) (*entities.Admin, error)
	GetAdminByID(id uuid.UUID) (*entities.Admin, error)
}

type RepositoryAdminImpl struct {
	db database.Database
}

func NewRepositoryAdmin(db database.Database) *RepositoryAdminImpl  {
	return &RepositoryAdminImpl{db : db}
}

func (r *RepositoryAdminImpl) GetAdmin() (bool, error) {
    query := `SELECT COUNT(*) FROM admin`
    
    var count int
    err := r.db.Get(&count, query)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (r *RepositoryAdminImpl) CreateAdmin(admin *entities.Admin) error {
    query := `
        INSERT INTO admin (user_name, password, full_name)
        VALUES (:user_name, :password, :full_name)
    `

    _, err := r.db.NamedExec(query, admin)
    return err
}



func (r *RepositoryAdminImpl) GetAdminByUserName(userName string) (*entities.Admin, error) {
	query := `SELECT * FROM admin WHERE user_name = ?`
	
	var admin entities.Admin
	err := r.db.Get(&admin, query, userName)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}



func (r *RepositoryAdminImpl) GetAdminByID(id uuid.UUID) (*entities.Admin, error) {
	query := `SELECT * FROM admin WHERE id = ?`
	
	var admin entities.Admin
	err := r.db.Get(&admin, query, id)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}
