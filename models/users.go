package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Mail                  string `gorm:"uniqueIndex"`
	Active                bool
	LoginToken            string `gorm:"unique"`
	MailVeryficationToken string `gorm:"unique"`
}
