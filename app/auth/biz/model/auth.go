package model

import (
	"context"

	"gorm.io/gorm"
)

type Token struct {
	CreatedAt      int64
	ExpirationTime int64  `gorm:"not null"`
	ID             uint   `gorm:"primary_key;AUTO_INCREMENT"`
	UserID         string `gorm:"index;not null"`
	Value          string `gorm:"unique;not null"`
}

func (t Token) TableName() string {
	return "token"
}

func (t *Token) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(t).Error
}
