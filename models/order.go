package models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	ProductID uint
	OrderID   uint
	Quantity  uint64
	Total     float64
}

type Order struct {
	gorm.Model
	OrderProducts []OrderProduct
	Total         float64
}

type OrderProductSerializer struct {
	ID        uint
	ProductID uint
	Quantity  uint64
	Total     float64
}
