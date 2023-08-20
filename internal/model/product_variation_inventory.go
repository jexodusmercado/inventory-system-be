package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductVariationInventory struct {
	ID 					uuid.UUID 		`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	ProductVariationID 	uuid.UUID 		`gorm:"type: uuid; not null" json:"product_variation_id"`
	AvailableStock 		int 			`gorm:"type: int; not null" json:"available_stock"`
	CreatedAt 			time.Time 		`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 			time.Time 		`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 			gorm.DeletedAt 	`gorm:"index" json:"deleted_at"`
}