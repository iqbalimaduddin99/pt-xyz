package repository

import (
	"pt-xyz/internal/entities"
	"github.com/jmoiron/sqlx"
)

type RepositoryTransactionProduct interface {
	CreateTransactionProduct(tx *sqlx.DB,transactionProduct *entities.TransactionProduct) error
}

type RepositoryTransactionProductImpl struct {
	db *sqlx.DB
}

func NewRepositoryTransactionProduct() *RepositoryTransactionProductImpl {
	return &RepositoryTransactionProductImpl{}
}

func (r *RepositoryTransactionProductImpl) CreateTransactionProduct(tx *sqlx.DB,transactionProduct *entities.TransactionProduct) error {
	query := `
		INSERT INTO transaction_product (
			transaction_id, company_id, product_company_id, otr, asset_name, price
		)
		VALUES (
			:transaction_id, :company_id, :product_company_id, :otr, :asset_name, :price
		)
	`
 	_, err := tx.NamedExec(query, transactionProduct)
	return err
}
