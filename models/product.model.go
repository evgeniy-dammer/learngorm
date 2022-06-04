package models

import (
	"time"

	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type ProductModel struct{}

//FindAll returns entities list
func (productModel ProductModel) FindAll(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Find(&products)

	return products, nil
}

//FindProductsWithConditions returns entities list with conditions
func (productModel ProductModel) FindProductsWithConditions(db *gorm.DB, status bool, min float64, max float64) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("status = ? AND price >= ? AND price <= ?", status, min, max).Find(&products)

	return products, nil
}

//FindProductWithBetween returns entities list with BETWEEN condition
func (productModel ProductModel) FindProductsWithBetween(db *gorm.DB, min float64, max float64) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("price BETWEEN ? AND ?", min, max).Find(&products)

	return products, nil
}

//FindProductStartsWith returns entities list starts with LIKE condition
func (productModel ProductModel) FindProductsStartsWith(db *gorm.DB, keyword string) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("name LIKE ?", keyword+"%").Find(&products)

	return products, nil
}

//FindProductEndsWith returns entities list ends with LIKE condition
func (productModel ProductModel) FindProductsEndsWith(db *gorm.DB, keyword string) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("name LIKE ?", "%"+keyword).Find(&products)

	return products, nil
}

//FindProductContains returns entities list contains with LIKE condition
func (productModel ProductModel) FindProductsContains(db *gorm.DB, keyword string) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("name LIKE ?", "%"+keyword+"%").Find(&products)

	return products, nil
}

//FindProductsOrderByDesc returns entities list ordered by DESC
func (productModel ProductModel) FindProductsOrderByDesc(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Order("price DESC").Find(&products)

	return products, nil
}

//FindProductsOrderByAsc returns entities list ordered by ASC
func (productModel ProductModel) FindProductsOrderByAsc(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Order("price ASC").Find(&products)

	return products, nil
}

//FindProductsOrderByAndCondition returns entities list ordered by and condition
func (productModel ProductModel) FindProductsOrderByAndCondition(db *gorm.DB, status bool) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("status = ?", status).Order("price DESC").Find(&products)

	return products, nil
}

//FindProductsWithLimit returns entities list with LIMIT
func (productModel ProductModel) FindProductsWithLimit(db *gorm.DB, n int) ([]entities.Product, error) {
	var products []entities.Product

	db.Limit(n).Find(&products)

	return products, nil
}

//FindProductsOrderByWithLimit returns entities list ordered by with LIMIT
func (productModel ProductModel) FindProductsOrderByWithLimit(db *gorm.DB, n int) ([]entities.Product, error) {
	var products []entities.Product

	db.Order("price DESC").Limit(n).Find(&products)

	return products, nil
}

//FindProductsOrderByWithWhereAndLimit returns entities list ordered by with LIMIT and condition
func (productModel ProductModel) FindProductsOrderByWithWhereAndLimit(db *gorm.DB, status bool, n int) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("status = ?", status).Order("price DESC").Limit(n).Find(&products)

	return products, nil
}

//FindProductsByYearAndMonthAndDay returns entities list by year, month and day
func (productModel ProductModel) FindProductsByYearAndMonthAndDay(db *gorm.DB, year int, month int, day int) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("extract(year from created) = ? AND extract(month from created) = ? AND extract(day from created) = ?", year, month, day).Find(&products)

	return products, nil
}

//FindProductsByDate returns entities list by date
func (productModel ProductModel) FindProductsByDate(db *gorm.DB, date time.Time) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("created = ?", date).Find(&products)

	return products, nil
}

//FindProductsByDates returns entities list by dates
func (productModel ProductModel) FindProductsByDates(db *gorm.DB, startDate time.Time, endDate time.Time) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("created >= ? AND created <= ?", startDate, endDate).Find(&products)

	return products, nil
}

//FindProductById returns entity by id
func (productModel ProductModel) FindProductById(db *gorm.DB, id int) ([]entities.Product, error) {
	var products []entities.Product

	db.Where("id = ?", id).Find(&products)

	return products, nil
}

//FindProductWithSelect returns entities with select
func (productModel ProductModel) FindProductWithSelect(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Table("product").Select("*").Scan(&products)

	return products, nil
}

//FindProductWithSelectAndCondition returns entities with select and condition
func (productModel ProductModel) FindProductWithSelectAndCondition(db *gorm.DB, status bool) ([]entities.Product, error) {
	var products []entities.Product

	db.Table("product").Where("status = ?", status).Select("*").Scan(&products)

	return products, nil
}

//FindProductGroupBy returns entities list groupped by
func (productModel ProductModel) FindProductGroupBy(db *gorm.DB) ([]entities.ProductGroup, error) {
	var productGroups []entities.ProductGroup

	db.Table("product").Select("status, count(id) AS result1, sum(quantity) AS result2, min(price) AS result3, max(price) AS result4, avg(price) AS result5").Group("status").Scan(&productGroups)

	return productGroups, nil
}

//FindProductGroupBy returns entities list with HAVING
func (productModel ProductModel) FindProductWithHaving(db *gorm.DB) ([]entities.ProductGroup, error) {
	var productGroups []entities.ProductGroup

	db.Table("product").Select("status, count(id) AS result1, sum(quantity) AS result2, min(price) AS result3, max(price) AS result4, avg(price) AS result5").Group("status").Having("count(id) > ?", 2).Scan(&productGroups)

	return productGroups, nil
}

//FindSum returns SUM by quantity
func (productModel ProductModel) FindSum(db *gorm.DB) int64 {
	var result int64

	row := db.Table("product").Select("sum(quantity)").Row()
	row.Scan(&result)

	return result
}

//FindSumWithCondition returns SUM by quantity with conditions
func (productModel ProductModel) FindSumWithCondition(db *gorm.DB, status bool) int64 {
	var result int64

	row := db.Table("product").Where("status = ?", status).Select("sum(quantity)").Row()
	row.Scan(&result)

	return result
}

//FindSumWithCalculate returns SUM of calculate with conditions
func (productModel ProductModel) FindSumWithCalculate(db *gorm.DB, status bool) float64 {
	var result float64

	row := db.Table("product").Where("status = ?", status).Select("sum(price * quantity)").Row()
	row.Scan(&result)

	return result
}

//FindCount returns COUNT of entities
func (productModel ProductModel) FindCount(db *gorm.DB) int64 {
	var result int64

	db.Table("product").Count(&result)

	return result
}

//FindCountWithCondition returns COUNT of entities wit condition
func (productModel ProductModel) FindCountWithCondition(db *gorm.DB, status bool) int64 {
	var result int64

	db.Table("product").Where("status = ?", status).Count(&result)

	return result
}

//FindMin returns MIN of price
func (productModel ProductModel) FindMin(db *gorm.DB) float64 {
	var result float64

	row := db.Table("product").Select("min(price)").Row()
	row.Scan(&result)

	return result
}

//FindMinWithCondition returns MIN of price with condition
func (productModel ProductModel) FindMinWithCondition(db *gorm.DB, status bool) float64 {
	var result float64

	row := db.Table("product").Where("status = ?", status).Select("min(price)").Row()
	row.Scan(&result)

	return result
}

//FindMax returns MAX of price
func (productModel ProductModel) FindMax(db *gorm.DB) float64 {
	var result float64

	row := db.Table("product").Select("max(price)").Row()
	row.Scan(&result)

	return result
}

//FindMaxWithCondition returns MAX of price with condition
func (productModel ProductModel) FindMaxWithCondition(db *gorm.DB, status bool) float64 {
	var result float64

	row := db.Table("product").Where("status = ?", status).Select("max(price)").Row()
	row.Scan(&result)

	return result
}

//FindAvg returns AVG of price
func (productModel ProductModel) FindAvg(db *gorm.DB) float64 {
	var result float64

	row := db.Table("product").Select("avg(price)").Row()
	row.Scan(&result)

	return result
}

//FindAvgWithCondition returns AVG of price with condition
func (productModel ProductModel) FindAvgWithCondition(db *gorm.DB, status bool) float64 {
	var result float64

	row := db.Table("product").Where("status = ?", status).Select("avg(price)").Row()
	row.Scan(&result)

	return result
}

//FindAllStoredProc returns entities list using stored procedure
func (productModel ProductModel) FindAllStoredProc(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Raw("call sp_findAll()").Scan(&products)

	return products, nil
}

//CreateEntity creates entity
func (productModel ProductModel) CreateEntity(db *gorm.DB, product *entities.Product) error {
	db.Create(&product)

	return nil
}

//UpdateEntity updates entity
func (productModel ProductModel) UpdateEntity(db *gorm.DB, product *entities.Product) error {
	db.Save(&product)

	return nil
}

//FindProductByIdWithFirst returns entity by id
func (productModel ProductModel) FindProductByIdWithFirst(db *gorm.DB, id int) (entities.Product, error) {
	var product entities.Product

	db.Where("id = ?", id).First(&product)

	return product, nil
}

//DeleteEntity deletes entity
func (productModel ProductModel) DeleteEntity(db *gorm.DB, product entities.Product) error {
	db.Delete(&product)

	return nil
}

//FindAllWithStoredProcedure returns entities list with stored procedure
func (productModel ProductModel) FindAllWithStoredProcedure(db *gorm.DB) ([]entities.Product, error) {
	var products []entities.Product

	db.Raw("CALL sp_findAll()").Scan(&products)

	return products, nil
}

//FindAllWithStoredProcedure returns entities list with stored procedure
func (productModel ProductModel) FindByPricesWithStoredProcedure(db *gorm.DB, min, max float64) ([]entities.Product, error) {
	var products []entities.Product

	db.Raw("CALL findBetween(?, ?)", min, max).Scan(&products)

	return products, nil
}
