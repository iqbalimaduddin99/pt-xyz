package entities

import (
	"time"

	"github.com/google/uuid"
)

type LoanInstallment struct {
	Master
	TransactionID   uuid.UUID  `db:"transaction_id" json:"transactionId"`
	ConsumerID      uuid.UUID  `db:"consumer_id" json:"consumerId"`
	InstallmentAmount float64  `db:"installment_amount" json:"installmentAmount"`
	Tenor           time.Time  `db:"tenor" json:"tenor"`
	InterestRate    float64    `db:"interest_rate" json:"interestRate"`
}
