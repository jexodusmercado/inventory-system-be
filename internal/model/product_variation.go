package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductVariation struct {
	ID 				uuid.UUID 		`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	ProductID 		uuid.UUID 		`gorm:"type: uuid; not null" json:"product_id"`
	Unit 			string 			`gorm:"type: varchar(255); not null" json:"unit"`
	UnitValue 		int				`gorm:"type: int; not null" json:"unit_value"`
	CostPrice 		float64 		`gorm:"type: float; not null" json:"cost_price"`
	Price 			float64 		`gorm:"type: float; not null" json:"price"`
	Inventory		ProductVariationInventory `gorm:"foreignKey: ProductVariationID;" json:"inventory"`
	CreatedAt 		time.Time 		`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 		time.Time 		`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt 	`gorm:"index" json:"deleted_at"`
}