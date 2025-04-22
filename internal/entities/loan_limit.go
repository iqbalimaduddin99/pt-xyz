
package entities

import (

	"github.com/google/uuid"
)

type LoanLimit struct {
	Master
	ConsumerID    uuid.UUID  `db:"consumer_id" json:"consumerId"`
	LimitLoan     float64    `db:"limit_loan" json:"limitLoan"`
	LimitUsed     float64    `db:"limit_used" json:"limitUsed"`
	TenorAmount   int        `db:"tenor_amount" json:"tenorAmount"`
}
