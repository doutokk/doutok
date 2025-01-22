package model

type Token struct {
	ID             uint   `gorm:"primary_key;AUTO_INCREMENT"`
	UserID         string `gorm:"index;not null"`
	Value          string `gorm:"unique;not null"`
	ExpirationTime int64  `gorm:"not null"`
	CreatedAt      int64
}
