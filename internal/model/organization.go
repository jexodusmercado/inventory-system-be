package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID 			uuid.UUID 			`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	Name 		string 				`gorm:"type: varchar(255); not null" json:"name"`
	Description	string 				`gorm:"type: int; not null" json:"description"`
	Details		OrganizationDetail 	`gorm:"foreignKey:OrganizationID" json:"details"`
	Users 		[]User 				`gorm:"foreignKey:OrganizationID" json:"users"`
	Addresses   []Address 			`gorm:"foreignKey:BelongTo" json:"addresses"`
	CreatedAt 	time.Time 			`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 	time.Time 			`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 		`gorm:"index" json:"deleted_at"`
}

// NewOrganization and save to database
func (o *Organization) NewOrganization(db *gorm.DB) (*Organization, error) {
	err := db.Create(&o).Error
	if err != nil {
		return &Organization{}, err
	}
	return o, nil
}

// GetOrganizationByID get organization by id
func (o *Organization) GetOrganizationByID(db *gorm.DB, uid uuid.UUID) (*Organization, error) {
	err := db.Model(Organization{}).Where("id = ?", uid).Take(&o).Error
	if err != nil {
		return &Organization{}, err
	}
	return o, nil
}