package repository

import (
	"pt-xyz/internal/entities"

	"github.com/jmoiron/sqlx"
)

type RepositoryAdmin interface {
	CheckAdminExists() (bool, error)
	InsertAdmin(admin *entities.Admin) error
}

type RepositoryAdminImpl struct {
	db *sqlx.DB
}

func NewRepositoryAdmin(db *sqlx.DB) *RepositoryAdminImpl  {
	return &RepositoryAdminImpl{db : db}
}

func (r *RepositoryAdminImpl) CheckAdminExists() (bool, error) {
    query := `SELECT COUNT(*) FROM admin`
    
    var count int
    err := r.db.Get(&count, query)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (r *RepositoryAdminImpl) InsertAdmin(admin *entities.Admin) error {
    query := `
        INSERT INTO admin (user_name, password, full_name)
        VALUES (:user_name, :password, :full_name)
    `

    _, err := r.db.NamedExec(query, admin)
    return err
}
