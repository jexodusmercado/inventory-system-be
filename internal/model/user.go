package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID 							uuid.UUID		`gorm:"type: uuid; primaryKey; default: uuid_generate_v4()" json:"id"`
	OrganizationID				uuid.UUID 		`gorm:"type: uuid; not null" json:"organization_id"`
	Organization				Organization	`gorm:"foreignKey:OrganizationID" json:"organization"`
	Details						UserDetail		`gorm:"foreignKey:UserID" json:"details"`

	Email 						string 			`gorm:"type: varchar(255); not null" json:"email"`
	EmailConfirmedAt 			time.Time		`gorm:"type: timestamptz;" json:"email_confirmed_at"`
	EmailChangeToken 			string 			`gorm:"type: varchar(255);" json:"email_change_token"`
	EmailChangeSentAt 			time.Time 		`gorm:"type: timestamptz;" json:"email_change_sent_at"`
	EmailChangeConfirmStatus 	bool 			`gorm:"type: boolean; not null" json:"email_change_confirm_status"`

	Phone 						string			`gorm:"type: varchar(255);" json:"phone"`
	PhoneConfirmedAt 			time.Time 		`gorm:"type: timestamptz;" json:"phone_confirmed_at"`
	PhoneChangeToken 			string 			`gorm:"type: varchar(255);" json:"phone_change_token"`
	PhoneChangeSentAt 			time.Time 		`gorm:"type: timestamptz;" json:"phone_change_sent_at"`
	PhoneChangeConfirmStatus 	bool 			`gorm:"type: boolean; not null" json:"phone_change_confirm_status"`

	Password 					string 			`gorm:"type: varchar(255); not null" json:"password"`
	PasswordResetToken 			string 			`gorm:"type: varchar(255);" json:"password_reset_token"`
	PasswordResetSentAt 		time.Time 		`gorm:"type: timestamptz;" json:"password_reset_sent_at"`
	PasswordResetConfirmStatus 	bool 			`gorm:"type: boolean;" json:"password_reset_confirm_status"`

	LastSignInAt 				time.Time 		`gorm:"type: timestamptz;" json:"last_sign_in_at"`

	InvitedAt 					time.Time 		`gorm:"type: timestamptz;" json:"invited_at"`
	CreatedAt 					time.Time 		`gorm:"type: timestamptz;" json:"created_at"`
	UpdatedAt 					time.Time 		`gorm:"type: timestamptz;" json:"updated_at"`
	DeletedAt 					gorm.DeletedAt 	`gorm:"index" json:"deleted_at"`
}