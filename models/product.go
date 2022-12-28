package models

import (
	"AnarchyStock/database"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string         `json:"name"`
	Category string         `json:"category"`
	Price    float64        `json:"price"`
	Quantity uint64         `json:"quantity"`
	Orders   []OrderProduct `json:"orders"`
}

type ProductSerializer struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Quantity uint64  `json:"quantity"`
}

func CreateProduct(product ProductSerializer) {
	prod := Product{
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	database.DB.DB.Create(&prod)
}
func GetProductByID(pk uint) Product {
	var product Product
	database.DB.DB.First(&product, pk)
	return product
}
func GetAllProducts() []Product {
	var products []Product
	database.DB.DB.Order("Name").Find(&products)
	return products
}
func GetProductsCount() (count int64) {
	database.DB.DB.Model(&Product{}).Count(&count)
	return count
}
func CheckProductStock(pk uint, quantity uint64) bool {
	var product Product
	database.DB.DB.First(&product, pk)
	return product.Quantity >= quantity
}
func GetProductsPaginated(page uint) []Product {
	var products []Product
	res := database.DB.DB.Limit(database.PAGE_SIZE).Offset(int(page-1) * database.PAGE_SIZE).Find(&products)
	if res.RowsAffected == 0 {
		return nil
	}
	return products
}
