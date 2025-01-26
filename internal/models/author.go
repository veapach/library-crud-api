package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}
