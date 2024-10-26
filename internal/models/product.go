package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	MarketID uint    `json:"market_id"`
}
