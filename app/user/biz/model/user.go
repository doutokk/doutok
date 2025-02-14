package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Querier interface {
	// GetOneByEmail
	//
	// SELECT * FROM @@table WHERE email = @email LIMIT 1
	GetOneByEmail(email string) (*User, error)
}
