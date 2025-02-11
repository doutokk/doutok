package model

import "gorm.io/gorm"

// Order 订单主表
type Order struct {
	gorm.Model
	OrderID      string `gorm:"type:varchar(128);primary_key"` // 订单 ID
	UserID       uint32 `gorm:"not null;index"`                // 用户 ID
	UserCurrency string `gorm:"type:varchar(10);not null"`     // 用户货币类型
	Email        string `gorm:"type:varchar(255);not null"`    // 用户邮箱

	// 地址信息
	StreetAddress string `gorm:"type:varchar(255);not null"` // 街道地址
	City          string `gorm:"type:varchar(100);not null"` // 城市
	State         string `gorm:"type:varchar(100);not null"` // 州/省
	Country       string `gorm:"type:varchar(100);not null"` // 国家
	ZipCode       int32  `gorm:"not null"`                   // 邮政编码

	CreatedAt  int32 `gorm:"not null"`      // 创建时间
	PaidStatus bool  `gorm:"default:false"` // 支付状态
	//OrderItems []OrderItem `gorm:"foreignKey:OrderID"` // 订单商品列表
}

// TableName 指定表名
func (Order) TableName() string {
	return "orders"
}

// OrderItem 订单商品表
type OrderItem struct {
	gorm.Model
	OrderID   string  `gorm:"type:varchar(128);not null;index"` // 订单 ID
	ProductID uint32  `gorm:"not null"`                         // 商品 ID
	Quantity  int32   `gorm:"not null"`                         // 商品数量
	Cost      float64 `gorm:"type:decimal(10,2);not null"`      // 商品成本
}

// TableName 指定表名
func (OrderItem) TableName() string {
	return "order_items"
}

type Querier interface {
	// GetByUserId 根据用户 ID 获取订单列表
	//
	// SELECT * FROM orders WHERE user_id = @userId AND deleted_at is null
	GetByUserId(userId uint32) ([]*Order, error)
}
