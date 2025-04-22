package repository

import (
	"fmt"
	"pt-xyz/internal/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RepositoryLoanLimit interface {
	GetLoanLimitByID(tx *sqlx.DB,consumerID uuid.UUID) (*entities.LoanLimit, error)
}

type RepositoryLoanLimitImpl struct {
	db *sqlx.DB
}

func NewRepositoryLoanLimit() *RepositoryLoanLimitImpl {
	return &RepositoryLoanLimitImpl{}
}

func (r *RepositoryLoanLimitImpl) GetLoanLimitByID(tx *sqlx.DB,consumerID uuid.UUID)  (*entities.LoanLimit, error) {
	
	fmt.Println("jaut", consumerID)
	query := `SELECT * FROM loan_limit WHERE consumer_id = ?`
	
	var loanLimit entities.LoanLimit
	err := tx.Get(&loanLimit, query, consumerID)
	if err != nil {
		return nil, err
	}

	fmt.Println("sss")
	return &loanLimit, nil
}
