package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Step struct {
	ID uuid.UUID `gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	StepName string `gorm:"type: varchar(255); not null" json:"step_name"`
	StepDescription string `gorm:"type: text; not null" json:"step_description"`
	StepNumber int `gorm:"type: int; not null" json:"step_number"`
	RecipeID uuid.UUID `gorm:"type: uuid; not null" json:"recipe_id"`
	CreatedAt time.Time `gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt time.Time `gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type: timestamptz;" json:"deleted_at"`
}