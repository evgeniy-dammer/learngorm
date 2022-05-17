package models

import (
	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type ProductModel struct{}

//FindAll returns entities list from Database
func (productModel ProductModel) FindAll(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Find(&products)

	return products, nil
}

//FindProductsWithConditions returns entities list with conditions from Database
func (productModel ProductModel) FindProductsWithConditions(db *gorm.DB, status bool, min float64, max float64) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("status = ? AND price >= ? AND price <= ?", status, min, max).Find(&products)

	return products, nil
}

//FindProductWithBetween returns entities list with BETWEEN condition from Database
func (productModel ProductModel) FindProductsWithBetween(db *gorm.DB, min float64, max float64) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("price BETWEEN ? AND ?", min, max).Find(&products)

	return products, nil
}
