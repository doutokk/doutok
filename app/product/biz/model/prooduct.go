package model

type Product struct {
	ID          uint              `gorm:"primary_key"`
	Name        string            `gorm:"type:varchar(128) not null"`
	Description string            `gorm:"type:varchar(256) not null"`
	Picture     string            `gorm:"type:varchar(256) not null"`
	Price       float32           `gorm:"type:decimal(10,2) not null"`
	Categories  []ProductCategory `gorm:"many2many:product_categories;"`
}

type ProductCategory struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"unique;not null"`
}
