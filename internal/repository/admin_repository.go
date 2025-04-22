package repository

import (
	"pt-xyz/internal/entities"

	"github.com/jmoiron/sqlx"
)

type RepositoryAdmin interface {
	GetAdmin() (bool, error)
	CreateAdmin(admin *entities.Admin) error
	GetAdminByUserName(userName string) (*entities.Admin, error)
}

type RepositoryAdminImpl struct {
	db *sqlx.DB
}

func NewRepositoryAdmin(db *sqlx.DB) *RepositoryAdminImpl  {
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
