package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	ID uuid.UUID 				`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	Url string 					`gorm:"type: varchar(255); not null" json:"url"`
	CreatedAt time.Time 		`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt time.Time 		`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt gorm.DeletedAt 	`gorm:"index"`
}