package model

import (
	"time"

	"github.com/google/uuid"
)

type Identity struct {
	ID           uuid.UUID	`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID	`gorm:"type: uuid; not null" json:"user_id"`
	Provider     string		`gorm:"type: varchar(255); not null" json:"provider"`
	LastSignInAt time.Time	`gorm:"type: timestamptz;" json:"last_sign_in_at"`
	CreatedAt    time.Time	`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt    time.Time	`gorm:"type: timestamptz;" json:"updated_at"`
}