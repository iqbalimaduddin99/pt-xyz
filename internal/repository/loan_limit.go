package repository

import (
	"database/sql"
	"pt-xyz/configs/database"
	"pt-xyz/internal/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RepositoryLoanLimit interface {
	GetLoanLimitByID(tx database.Database,consumerID uuid.UUID) (*entities.LoanLimit, error)
	CreateLoanLimit(limit *entities.LoanLimit) error
}

type RepositoryLoanLimitImpl struct {
	db *sqlx.DB
}

func NewRepositoryLoanLimit(db *sqlx.DB) *RepositoryLoanLimitImpl {
	return &RepositoryLoanLimitImpl{db: db}
}

func (r *RepositoryLoanLimitImpl) GetLoanLimitByID(tx database.Database,consumerID uuid.UUID) (*entities.LoanLimit, error) {

	query := `SELECT * FROM loan_limit WHERE consumer_id = ?`
	
	var loanLimit entities.LoanLimit
	err := tx.Get(&loanLimit, query, consumerID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if loanLimit.ConsumerID == uuid.Nil {
		return nil, nil
	}

	return &loanLimit, nil
}



func (r *RepositoryLoanLimitImpl) CreateLoanLimit(limit *entities.LoanLimit) error {
	query := `
		INSERT INTO loan_limit (
			consumer_id,
			limit_loan,
			limit_used,
			tenor_amount
		)
		VALUES (
			:consumer_id,
			:limit_loan,
			:limit_used,
			:tenor_amount
		)
	`

	_, err := r.db.NamedExec(query, limit)
	return err
}
