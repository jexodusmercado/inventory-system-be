package model

import (
	"github.com/google/uuid"
)

type UserDetail struct {
	ID 				uuid.UUID	`gorm:"type: uuid; primaryKey;" json:"id"`
	FirstName 		string 		`gorm:"type: varchar(255); not null" json:"first_name"`
	LastName 		string		`gorm:"type: varchar(255); not null" json:"last_name"`
	Avatar 			string		`gorm:"type: varchar(255); not null" json:"avatar"`
	Addresses 		[]Address	`gorm:"foreignKey:BelongTo" json:"addresses"`
}