package model

// Order 订单主表
type Order struct {
	OrderID      string `gorm:"column:order_id;type:varchar(64);primary_key"`   // 订单 ID
	UserID       uint32 `gorm:"column:user_id;not null;index"`                  // 用户 ID
	UserCurrency string `gorm:"column:user_currency;type:varchar(10);not null"` // 用户货币类型
	Email        string `gorm:"column:email;type:varchar(255);not null"`        // 用户邮箱

	// 地址信息
	StreetAddress string `gorm:"column:street_address;type:varchar(255);not null"` // 街道地址
	City          string `gorm:"column:city;type:varchar(100);not null"`           // 城市
	State         string `gorm:"column:state;type:varchar(100);not null"`          // 州/省
	Country       string `gorm:"column:country;type:varchar(100);not null"`        // 国家
	ZipCode       int32  `gorm:"column:zip_code;not null"`                         // 邮政编码

	CreatedAt  int32       `gorm:"column:created_at;not null"`       // 创建时间
	PaidStatus bool        `gorm:"column:paid_status;default:false"` // 支付状态
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`               // 订单商品列表
}

// TableName 指定表名
func (Order) TableName() string {
	return "orders"
}

// OrderItem 订单商品表
type OrderItem struct {
	ID        int64   `gorm:"column:id;primary_key;auto_increment"`            // 自增主键
	OrderID   string  `gorm:"column:order_id;type:varchar(64);not null;index"` // 订单 ID
	ProductID uint32  `gorm:"column:product_id;not null"`                      // 商品 ID
	Quantity  int32   `gorm:"column:quantity;not null"`                        // 商品数量
	Cost      float64 `gorm:"column:cost;type:decimal(10,2);not null"`         // 商品成本
}

// TableName 指定表名
func (OrderItem) TableName() string {
	return "order_items"
}

type Querier interface {
	// GetByUserId 根据用户ID获取订单列表
	//
	// SELECT * FROM orders WHERE user_id = @userId AND deleted_at is null
	GetByUserId(userId uint32) ([]*Order, error)
}
