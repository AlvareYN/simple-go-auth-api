package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Role     string `gorm:"not null"`
	Hash     string `gorm:"not null" json:"-"`
}
