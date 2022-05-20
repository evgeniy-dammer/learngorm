package controllers

import (
	"log"
	"time"

	"github.com/evgeniy-dammer/learngorm/config"
	"github.com/evgeniy-dammer/learngorm/models"
	"github.com/evgeniy-dammer/learngorm/output"
	"gorm.io/gorm"
)

var productModel models.ProductModel

var (
	db    *gorm.DB
	errDB error
)

func init() {
	db, errDB = config.GetDB()

	if errDB != nil {
		log.Fatal(errDB)
	}
}

//Get entities list from Database
func ProductsList() {
	productsList, err := productModel.FindAll(db)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsList, "List of all products:")
}

//Get entities list with conditions from Database
func ProductsListWithCondition() {
	productsListWithCondition, err := productModel.FindProductsWithConditions(db, true, 5, 15)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListWithCondition, "List of products with condition:")
}

//Get entities list with BETWEEN conditions from Database
func ProductsListWithBetween() {
	productsListWithBetween, err := productModel.FindProductsWithBetween(db, 5, 15)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListWithBetween, "List of products with BETWEEN condition:")
}

//Get entities list starts with LIKE conditions from Database
func ProductsListStartsWith() {
	productsListStartsWith, err := productModel.FindProductsStartsWith(db, "Mo")

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListStartsWith, "List of products starts with LIKE condition:")
}

//Get entities list ends with LIKE conditions from Database
func ProductsListEndsWith() {
	productsListEndsWith, err := productModel.FindProductsEndsWith(db, "vi 1")

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListEndsWith, "List of products ends with LIKE condition:")
}

//Get entities list contains with LIKE conditions from Database
func ProductsListContains() {
	productsListContains, err := productModel.FindProductsContains(db, "bi")

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListContains, "List of products contains with LIKE condition:")
}

//Get entities list ordered by DESC from Database
func ProductsListOrderByDesc() {
	productsListOrderByDesc, err := productModel.FindProductsOrderByDesc(db)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListOrderByDesc, "List of products ordered by DESC:")
}

//Get entities list ordered by ASC from Database
func ProductsListOrderByAsc() {
	productsListOrderByAsc, err := productModel.FindProductsOrderByAsc(db)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListOrderByAsc, "List of products ordered by ASC:")
}

//Get entities list ordered by ASC from Database
func ProductsListOrderByAndCondition() {
	productsListOrderByAndCondition, err := productModel.FindProductsOrderByAndCondition(db, true)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListOrderByAndCondition, "List of products ordered by DESC and with condition:")
}

//Get entities list with LIMIT from Database
func ProductsListWithLimit() {
	productsListWithLimit, err := productModel.FindProductsWithLimit(db, 2)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListWithLimit, "List of products with LIMIT:")
}

//Get entities list ordered by with LIMIT from Database
func ProductsListOrderByWithLimit() {
	productsListOrderByWithLimit, err := productModel.FindProductsOrderByWithLimit(db, 2)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListOrderByWithLimit, "List of products ordered by with LIMIT:")
}

//Get entities list ordered by with LIMIT and condition from Database
func ProductsListOrderByWithWhereAndLimit() {
	productsListOrderByWithWhereAndLimit, err := productModel.FindProductsOrderByWithWhereAndLimit(db, true, 2)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListOrderByWithWhereAndLimit, "List of products ordered by with LIMIT and condition:")
}

//Get entities list by year, month and day from Database
func ProductsListByYearAndMonthAndDay() {
	productsListByYearAndMonthAndDay, err := productModel.FindProductsByYearAndMonthAndDay(db, 2019, 11, 8)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListByYearAndMonthAndDay, "List of products by year, month and day:")
}

//Get entities list by date from Database
func ProductsListByDate() {
	date, err := time.Parse("2006-01-02", "2019-07-10")

	if err != nil {
		log.Fatalln(err)
	}

	productsListByDate, err := productModel.FindProductsByDate(db, date)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListByDate, "List of products by date:")
}

//Get entities list by dates from Database
func ProductsListByDates() {
	startDate, err := time.Parse("2006-01-02", "2019-07-09")

	if err != nil {
		log.Fatalln(err)
	}

	endDate, err := time.Parse("2006-01-02", "2019-07-15")

	if err != nil {
		log.Fatalln(err)
	}

	productsListByDates, err := productModel.FindProductsByDates(db, startDate, endDate)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productsListByDates, "List of products by dates:")
}

//Get entity by id from Database
func ProductById() {
	productById, err := productModel.FindProductById(db, 2)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productById, "Product by ID:")
}

//Get entities list with select from Database
func ProductListWithSelect() {
	productListWithSelect, err := productModel.FindProductWithSelect(db)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productListWithSelect, "List of products wit select:")
}

//Get entities list with select and condition from Database
func ProductListWithSelectAndCondition() {
	productListWithSelectAndCondition, err := productModel.FindProductWithSelectAndCondition(db, true)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productListWithSelectAndCondition, "List of products with select and condition:")
}

//Get entities list with GROUP BY from Database
func ProductListWithGroupBy() {
	productListWithGroupBy, err := productModel.FindProductGroupBy(db)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productListWithGroupBy, "List of products with GROUP BY:")
}

//Get entities list with HAVING from Database
func ProductListWithHaving() {
	productListWithHaving, err := productModel.FindProductWithHaving(db)

	if err != nil {
		log.Fatalln(err)
	}

	output.Output(productListWithHaving, "List of products with HAVING:")
}
