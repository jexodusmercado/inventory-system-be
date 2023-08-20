package model

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID       	uuid.UUID `gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	BelongTo 	uuid.UUID `gorm:"type: uuid; not null" json:"belong_to"`
	Address1 	string    `gorm:"type: varchar(255); not null" json:"address_1"`
	Address2 	string    `gorm:"type: varchar(255);" json:"address_2"`
	City     	string    `gorm:"type: varchar(255); not null" json:"city"`
	State    	string    `gorm:"type: varchar(255); not null" json:"state"`
	Zip      	string    `gorm:"type: varchar(255); not null" json:"zip"`
	Country  	string    `gorm:"type: varchar(255); not null" json:"country"`
	IsDefault	bool	  `gorm:"type: boolean; default: false; not null;" json:"is_default"`
	CreatedAt 	time.Time `gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 	time.Time `gorm:"index" json:"deleted_at"`
}