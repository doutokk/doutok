package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(255);not null;index"`
	HashedPassword string `gorm:"type:varchar(255);not null"`
}

// Querier is the interface for the query, you can implement it with your own query logic
type Querier interface {
	// GetByUserId get user by user id
	//
	// SELECT * FROM @@table WHERE user_id = @userId and deleted_at is null
	GetByUserId(userId uint32) ([]*User, error)
}
