package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Master struct {
	ID        uuid.UUID  		 `db:"id" json:"id"`
	CreatedAt time.Time  		 `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  		 `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time 		 `db:"deleted_at" json:"deletedAt"`
	CreatedBy sql.NullString     `db:"created_by" json:"createdBy"`
	UpdatedBy sql.NullString     `db:"updated_by" json:"updatedBy"`
	DeletedBy sql.NullString     `db:"deleted_by" json:"deletedBy"`
}
