package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Mail                  string `gorm:"uniqueIndex" json:"mail"`
	Active                bool
	LoginToken            string `gorm:"unique" json:"login_token"`
	MailVeryficationToken string `gorm:"unique" json:"mail_veryfication_token"`
}
