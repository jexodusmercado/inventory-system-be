package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID 				uuid.UUID 			`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	OrganizationID	uuid.UUID 			`gorm:"type: uuid; not null" json:"organization_id"`
	Name 			string 				`gorm:"type: varchar(255); not null" json:"name"`
	Description		string 				`gorm:"type: int; not null" json:"description"`
	Product			[]*Product 			`gorm:"many2many:product_category;"`
	CreatedAt 		time.Time 			`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 		time.Time 			`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt 		`gorm:"index" json:"deleted_at"`
}