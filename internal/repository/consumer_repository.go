package repository

import (
	"pt-xyz/internal/entities"
	"github.com/jmoiron/sqlx"
)

type RepositoryConsumer interface {
	GetConsumerById(id int) (*entities.Consumer, error)
	GetConsumerByUserName(userName string) (*entities.Consumer, error)
	GetConsumerByKTP(KTP string) (*entities.Consumer, error)
	CreateConsumer(consumer *entities.ReqConsumer) error
}

type RepositoryConsumerImpl struct {
	db *sqlx.DB
}

func NewRepositoryConsumer(db *sqlx.DB) *RepositoryConsumerImpl {
	return &RepositoryConsumerImpl{db: db}
}

func (r *RepositoryConsumerImpl) GetConsumerByKTP(KTP string) (*entities.Consumer, error) {
	query := `SELECT * FROM consumer WHERE ktp = ?`
	
	var consumer entities.Consumer
	err := r.db.Get(&consumer, query, KTP)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (r *RepositoryConsumerImpl) GetConsumerById(id int) (*entities.Consumer, error) {
	query := `SELECT * FROM consumer WHERE id = ?`
	
	var consumer entities.Consumer
	err := r.db.Get(&consumer, query, id)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (r *RepositoryConsumerImpl) GetConsumerByUserName(userName string) (*entities.Consumer, error) {
	query := `SELECT * FROM consumer WHERE user_name = ?`
	
	var consumer entities.Consumer
	err := r.db.Get(&consumer, query, userName)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (r *RepositoryConsumerImpl) CreateConsumer(consumer *entities.ReqConsumer) error {
	query := `
		INSERT INTO consumer (KTP, user_name, password, full_name, legal_name, born_location, born_date, photo_KTP, selfie_photo, salary)
		VALUES (:KTP, :user_name, :password, :full_name, :legal_name, :born_location, :born_date, :photo_KTP, :selfie_photo, :salary)
	`

	_, err := r.db.NamedExec(query, consumer)
	return err
}
