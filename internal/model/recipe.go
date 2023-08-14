package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	ID uuid.UUID `gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	Name string `gorm:"type: varchar(255); not null" json:"name"`
	Description string `gorm:"type: varchar(255); not null" json:"description"`
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients; constraint: OnUpdate:CASCADE,OnDelete:CASCADE;" json:"ingredients"`
	Steps []Step `gorm:"constraint: OnUpdate:CASCADE,OnDelete:CASCADE;" json:"steps"`
	Images []Image `gorm:"constraint: OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images"`
	CreatedAt time.Time `gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt time.Time `gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type: timestamptz;" json:"deleted_at"`
}