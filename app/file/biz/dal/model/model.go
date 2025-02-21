package model

import (
	"gorm.io/gorm"
)

// rename to create your own model
type File struct {
	gorm.Model
	UserId   uint32 `gorm:"type:int(11);not null;index"`
	FileName string `gorm:"type:varchar(255);not null"`
}

// Querier is the interface for the query, you can implement it with your own query logic
type Querier interface {
	// GetByUserId get user by user id
	//
	// SELECT * FROM @@table WHERE user_id = @userId and deleted_at is null
	GetByUserId(userId uint32) ([]*File, error)
}
