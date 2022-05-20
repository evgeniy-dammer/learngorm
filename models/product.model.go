package models

import (
	"time"

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

//FindProductStartsWith returns entities list starts with LIKE condition from Database
func (productModel ProductModel) FindProductsStartsWith(db *gorm.DB, keyword string) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("name LIKE ?", keyword+"%").Find(&products)

	return products, nil
}

//FindProductEndsWith returns entities list ends with LIKE condition from Database
func (productModel ProductModel) FindProductsEndsWith(db *gorm.DB, keyword string) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("name LIKE ?", "%"+keyword).Find(&products)

	return products, nil
}

//FindProductContains returns entities list contains with LIKE condition from Database
func (productModel ProductModel) FindProductsContains(db *gorm.DB, keyword string) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("name LIKE ?", "%"+keyword+"%").Find(&products)

	return products, nil
}

//FindProductsOrderByDesc returns entities list ordered by DESC from Database
func (productModel ProductModel) FindProductsOrderByDesc(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Order("price DESC").Find(&products)

	return products, nil
}

//FindProductsOrderByAsc returns entities list ordered by ASC from Database
func (productModel ProductModel) FindProductsOrderByAsc(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Order("price ASC").Find(&products)

	return products, nil
}

//FindProductsOrderByAndCondition returns entities list ordered by and condition from Database
func (productModel ProductModel) FindProductsOrderByAndCondition(db *gorm.DB, status bool) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("status = ?", status).Order("price DESC").Find(&products)

	return products, nil
}

//FindProductsWithLimit returns entities list with LIMIT from Database
func (productModel ProductModel) FindProductsWithLimit(db *gorm.DB, n int) ([]entities.Product, error) {
	var products []entities.Product

	db.Limit(n).Find(&products)

	return products, nil
}

//FindProductsOrderByWithLimit returns entities list ordered by with LIMIT from Database
func (productModel ProductModel) FindProductsOrderByWithLimit(db *gorm.DB, n int) ([]entities.Product, error) {
	var products []entities.Product

	db.Order("price DESC").Limit(n).Find(&products)

	return products, nil
}

//FindProductsOrderByWithWhereAndLimit returns entities list ordered by with LIMIT and condition from Database
func (productModel ProductModel) FindProductsOrderByWithWhereAndLimit(db *gorm.DB, status bool, n int) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("status = ?", status).Order("price DESC").Limit(n).Find(&products)

	return products, nil
}

//FindProductsByYearAndMonthAndDay returns entities list by year, month and day from Database
func (productModel ProductModel) FindProductsByYearAndMonthAndDay(db *gorm.DB, year int, month int, day int) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("extract(year from created) = ? AND extract(month from created) = ? AND extract(day from created) = ?", year, month, day).Find(&products)

	return products, nil
}

//FindProductsByDate returns entities list by date from Database
func (productModel ProductModel) FindProductsByDate(db *gorm.DB, date time.Time) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("created = ?", date).Find(&products)

	return products, nil
}

//FindProductsByDates returns entities list by dates from Database
func (productModel ProductModel) FindProductsByDates(db *gorm.DB, startDate time.Time, endDate time.Time) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("created >= ? AND created <= ?", startDate, endDate).Find(&products)

	return products, nil
}

//FindProductById returns entity by id from Database
func (productModel ProductModel) FindProductById(db *gorm.DB, id int) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("id = ?", id).Find(&products)

	return products, nil
}

//FindProductWithSelect returns entities with select from Database
func (productModel ProductModel) FindProductWithSelect(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Table("product").Select("*").Scan(&products)

	return products, nil
}

//FindProductWithSelectAndCondition returns entities with select and condition from Database
func (productModel ProductModel) FindProductWithSelectAndCondition(db *gorm.DB, status bool) ([]entities.Product, error) {
	var products []entities.Product

	db.Table("product").Where("status = ?", status).Select("*").Scan(&products)

	return products, nil
}

//FindProductGroupBy returns entities list groupped by from Database
func (productModel ProductModel) FindProductGroupBy(db *gorm.DB) ([]entities.ProductGroup, error) {
	var productGroups []entities.ProductGroup

	db.Table("product").Select("status, count(id) AS result1, sum(quantity) AS result2, min(price) AS result3, max(price) AS result4, avg(price) AS result5").Group("status").Scan(&productGroups)

	return productGroups, nil
}

//FindProductGroupBy returns entities list with HAVING from Database
func (productModel ProductModel) FindProductWithHaving(db *gorm.DB) ([]entities.ProductGroup, error) {
	var productGroups []entities.ProductGroup

	db.Table("product").Select("status, count(id) AS result1, sum(quantity) AS result2, min(price) AS result3, max(price) AS result4, avg(price) AS result5").Group("status").Having("count(id) > ?", 2).Scan(&productGroups)

	return productGroups, nil
}
