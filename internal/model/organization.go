package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID 			uuid.UUID 		`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	Name 		string 			`gorm:"type: varchar(255); not null" json:"name"`
	Description	string 			`gorm:"type: int; not null" json:"description"`
	Users 		[]User 			`gorm:"foreignKey:OrganizationID" json:"users"`
	Addresses   []Address 		`gorm:"foreignKey:OrganizationID" json:"addresses"`
	CreatedAt 	time.Time 		`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 	time.Time 		`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted_at"`
}