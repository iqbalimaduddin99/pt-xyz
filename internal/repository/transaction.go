package repository

import (
	"encoding/binary"
	"fmt"
	"pt-xyz/internal/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RepositoryTransaction interface {
	CreateTransaction(tx *sqlx.DB,transaction *entities.TransactionTable) (uuid.UUID, error)
}

type RepositoryTransactionImpl struct {
	db *sqlx.Tx
}

func NewRepositoryTransaction() *RepositoryTransactionImpl {
	return &RepositoryTransactionImpl{}
}

func (r *RepositoryTransactionImpl) CreateTransaction(tx *sqlx.DB,transaction *entities.TransactionTable) (uuid.UUID, error) {
	
	fmt.Println(transaction)
	fmt.Println(transaction.CompanyID)
	query := `
    INSERT INTO transaction_table (
        consumer_id, company_id, external_transaction_id, company_name,
        company_category, contact_number, admin_fee, total_price
    ) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

_, err := tx.Exec(query,
    transaction.ConsumerID,
    transaction.CompanyID,
    transaction.ExternalTransactionID,
    transaction.CompanyName,
    transaction.CompanyCategory,
    transaction.ContactNumber,
    transaction.AdminFee,
    transaction.TotalPrice,
)

if err != nil {
    return uuid.UUID{}, err
}

var transactionID []uint8
query = `
		SELECT id 
		FROM transaction_table
		WHERE consumer_id = ?
		ORDER BY created_at DESC 
		LIMIT 1;
		`

err = tx.QueryRow(query, transaction.ConsumerID).Scan(&transactionID)

if err != nil {
return uuid.UUID{}, err 
}

transactionIDStr := string(transactionID)

parsedUUID, err := uuid.Parse(transactionIDStr)
if err != nil {
    return uuid.UUID{}, fmt.Errorf("failed to parse UUID: %v", err)
}

fmt.Println("Parsed Transaction UUID:", parsedUUID)

	// transactionID := uint64ToUUID(lastInsertID)
	// fmt.Println("isinya", transactionID)
	fmt.Println("scs", transaction.ConsumerID)
	return parsedUUID, nil
}

	
func uint64ToUUID(id uint64) uuid.UUID {
	var uuidBytes [16]byte
	binary.BigEndian.PutUint64(uuidBytes[:8], id) 
	return uuid.UUID(uuidBytes)
}

