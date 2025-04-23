package repository

import (
	"pt-xyz/internal/entities"

	"github.com/jmoiron/sqlx"
)

type RepositoryLoanInstallment interface {
	CreateLoanInstallment(tx *sqlx.Tx,LoanInstallment *entities.LoanInstallment) error
}

type RepositoryLoanInstallmentImpl struct {
	db *sqlx.Tx
}

func NewRepositoryLoanInstallment() *RepositoryLoanInstallmentImpl {
	return &RepositoryLoanInstallmentImpl{}
}

func (r *RepositoryLoanInstallmentImpl) CreateLoanInstallment(tx *sqlx.Tx,LoanInstallment *entities.LoanInstallment) error {
	query := `
		INSERT INTO loan_installment (
			transaction_id, consumer_id, installment_amount, tenor, interest_rate
		)
		VALUES (
			:transaction_id, :consumer_id, :installment_amount, :tenor, :interest_rate
		)
	`
 	_, err := tx.NamedExec(query, LoanInstallment)//SQL Injection
	return err
}
