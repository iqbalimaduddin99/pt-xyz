package repository

import (
	"pt-xyz/internal/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RepositoryMasterProductXYZ interface {
	GetMasterProductForTransactionById(tx *sqlx.DB,id uuid.UUID)  (*entities.MasterProductPtXyz, error)
}

type RepositoryMasterProductXYZImpl struct {
	db *sqlx.Tx
}

func NewRepositoryMasterProductXYZ() *RepositoryMasterProductXYZImpl {
	return &RepositoryMasterProductXYZImpl{}
}

func (r *RepositoryMasterProductXYZImpl) GetMasterProductForTransactionById(tx *sqlx.DB,id uuid.UUID) (*entities.MasterProductPtXyz, error) {
	
	query := `SELECT * FROM master_product_pt_xyz WHERE id = ? FOR UPDATE`
	var masterProductPtXyz entities.MasterProductPtXyz
	err := tx.Get(&masterProductPtXyz, query, id)
	if err != nil {
		return nil, err
	}


	return &masterProductPtXyz, nil
}
