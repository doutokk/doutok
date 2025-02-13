package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string            `gorm:"type:varchar(128) not null"`
	Description string            `gorm:"type:varchar(256) not null"`
	Picture     string            `gorm:"type:varchar(256) not null"`
	Price       float32           `gorm:"type:decimal(10,2) not null"`
	Categories  []ProductCategory `gorm:"many2many:product_categories;"`
}

type ProductCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type Querier interface {
	// GetByUserId get user by user id
	//
	// SELECT * FROM @@table WHERE user_id = @userId and deleted_at is null
	GetByUserId(userId uint32) ([]*Product, error)
}
