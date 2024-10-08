package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"not null"`
	Email    string  `gorm:"not null;unique"`
	Role     string  `gorm:"not null"`
	Password string  `gorm:"not null"`
	Deposit  float64 `gorm:"not null"`
}
