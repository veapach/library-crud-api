package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string  `gorm:"not null"`
	AuthorID int     `gorm:"not null"`
	Author   Author  `gorm:"foreignKey:AuthorID"`
	Price    float64 `gorm:"not null"`
}
