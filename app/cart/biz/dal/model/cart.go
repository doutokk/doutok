package model

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UserId   uint `gorm:"type:int(11);not null;index"`
	ItemId   uint `gorm:"type:int(11);not null;index"`
	Quantity uint `gorm:"type:int(11);not null"`
}

type Querier interface {
}
