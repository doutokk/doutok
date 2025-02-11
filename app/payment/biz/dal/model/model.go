package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId uint32 `gorm:"type:int(11);not null;index"`
}

type Querier interface {
	// GetByUserId get user by user id
	//
	// SELECT * FROM @@table WHERE user_id = @userId and deleted_at is null
	GetByUserId(userId uint32) ([]*User, error)
}
