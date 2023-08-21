package storage

import (
	"github.com/jexodusmercado/inventory-system-be/internal/model"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Address{}, &model.Category{}, &model.Identity{},
		&model.Image{}, &model.Organization{}, &model.OrganizationDetail{},
		&model.Product{}, &model.ProductVariation{}, &model.ProductVariationInventory{},
		&model.User{}, &model.UserDetail{},
	)
}