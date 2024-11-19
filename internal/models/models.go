package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string    `gorm:"type:varchar(50);not null;unique" json:"username"`
	Email        string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password     string    `gorm:"type:varchar(255);not null" json:"password"`
	IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
	OTP          string    `gorm:"type:varchar(6)" json:"otp"`
	OTPCreatedAt time.Time `json:"otp_created_at"`
	Orders       []Order   `json:"orders,omitempty"` // One-to-Many relationship
	Cart         Cart      `json:"cart,omitempty"`   // One-to-One relationship
	Token        string    `json:"token"`
	IsActive     bool      `json:"is_active,omitempty"`
}

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(100);not null" json:"name"`
	Description string   `gorm:"type:text" json:"description"`
	Price       float64  `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int      `json:"stock"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"` // Belongs-To relationship
	MarketID    uint     `json:"market_id"`
	Market      Market   `gorm:"foreignKey:MarketID" json:"market,omitempty"` // Belongs-To relationship
}

type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	ParentID *uint     `json:"parent_id,omitempty"`                             // Nullable for top-level categories
	Parent   *Category `gorm:"foreignKey:ParentID" json:"parent,omitempty"`     // Self-referential relationship
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"` // One-to-Many relationship
}

type Order struct {
	gorm.Model
	UserID      uint    `json:"user_id"`
	User        User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TotalAmount float64 `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status      string  `gorm:"type:varchar(20)" json:"status"`
}

type Cart struct {
	gorm.Model
	UserID uint `gorm:"unique" json:"user_id"` // One-to-One relationship with User
}

type Market struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Address     string    `gorm:"type:varchar(255)" json:"address"`
	Phone       string    `gorm:"type:varchar(20)" json:"phone"`
	Email       string    `gorm:"type:varchar(100)" json:"email"`
	Products    []Product `gorm:"foreignKey:MarketID" json:"products,omitempty"` // One-to-Many relationship
}
