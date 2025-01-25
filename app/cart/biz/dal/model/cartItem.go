package model

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index"`
	ProductId uint32 `gorm:"type:int(11);not null;index"`
	Quantity  uint32 `gorm:"type:int(11);not null"`
}

type Querier interface {
	// GetByUserId get cart items by user id
	//
	// SELECT * FROM @@table WHERE user_id = @userId
	GetByUserId(userId uint32) ([]*CartItem, error)

	// GetByUserIdAndProductId get cart item by user id and item id
	//
	// SELECT * FROM @@table WHERE user_id = @userId AND product_id = @productId
	GetByUserIdAndProductId(userId, productId uint32) (*CartItem, error)
}
