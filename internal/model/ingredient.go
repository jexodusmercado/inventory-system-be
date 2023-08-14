package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ingredient struct {
	ID uuid.UUID `gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	Name string `gorm:"type: varchar(255); not null" json:"name"`
	Quantity int `gorm:"type: int; not null" json:"quantity"`
	Unit string `gorm:"type: varchar(255); not null" json:"unit"`
	RecipeID uuid.UUID `gorm:"type: uuid; not null" json:"recipe_id"`
	CreatedAt time.Time `gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt time.Time `gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type: timestamptz;" json:"deleted_at"`
}

func (i *Ingredient) TableName() string {
	return "ingredients"
}

type IngredientRepository interface {
	CreateIngredient(ingredient *Ingredient) error
	UpdateIngredient(ingredient *Ingredient) error
	DeleteIngredient(ingredient *Ingredient) error
	GetIngredientByID(id uuid.UUID) (*Ingredient, error)
	GetIngredientsByRecipeID(recipeID uuid.UUID) ([]*Ingredient, error)
}

type IngredientService interface {
	CreateIngredient(ingredient *Ingredient) error
	UpdateIngredient(ingredient *Ingredient) error
	DeleteIngredient(ingredient *Ingredient) error
	GetIngredientByID(id uuid.UUID) (*Ingredient, error)
	GetIngredientsByRecipeID(recipeID uuid.UUID) ([]*Ingredient, error)
}

type IngredientUsecase interface {
	CreateIngredient(ingredient *Ingredient) error
	UpdateIngredient(ingredient *Ingredient) error
	DeleteIngredient(ingredient *Ingredient) error
	GetIngredientByID(id uuid.UUID) (*Ingredient, error)
	GetIngredientsByRecipeID(recipeID uuid.UUID) ([]*Ingredient, error)
}

type IngredientRepositoryImpl struct {
	DB *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepositoryImpl {
	return &IngredientRepositoryImpl{DB: db}
}

func (r *IngredientRepositoryImpl) CreateIngredient(ingredient *Ingredient) error {
	return r.DB.Create(ingredient).Error
}

func (r *IngredientRepositoryImpl) UpdateIngredient(ingredient *Ingredient) error {
	return r.DB.Save(ingredient).Error
}

func (r *IngredientRepositoryImpl) DeleteIngredient(ingredient *Ingredient) error {
	return r.DB.Delete(ingredient).Error
}

func (r *IngredientRepositoryImpl) GetIngredientByID(id uuid.UUID) (*Ingredient, error) {
	var ingredient Ingredient
	if err := r.DB.Where("id = ?", id).First(&ingredient).Error; err != nil {
		return nil, err
	}
	return &ingredient, nil
}
