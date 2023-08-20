package model

import (
	"github.com/google/uuid"
)

type OrganizationDetail struct {
	ID 				uuid.UUID	`gorm:"type: uuid; primaryKey;" json:"id"`
	WebSite 		string		`gorm:"type: varchar(255); not null" json:"web_site"`
	Email 			string		`gorm:"type: varchar(255); not null" json:"email"`
	Phone 			string 		`gorm:"type: varchar(255); not null" json:"phone"`
	Address 		string 		`gorm:"type: varchar(255); not null" json:"address"`
	Avatar 			string 		`gorm:"type: varchar(255); not null" json:"avatar"`
}