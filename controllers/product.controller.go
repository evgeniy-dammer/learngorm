package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/evgeniy-dammer/learngorm/config"
	"github.com/evgeniy-dammer/learngorm/entities"
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

//Get entities list
func ProductsList() {
	productsList, err := productModel.FindAll(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsList, "List of all products:")
}

//Get entities list with conditions
func ProductsListWithCondition() {
	productsListWithCondition, err := productModel.FindProductsWithConditions(db, true, 5, 15)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListWithCondition, "List of products with condition:")
}

//Get entities list with BETWEEN conditions
func ProductsListWithBetween() {
	productsListWithBetween, err := productModel.FindProductsWithBetween(db, 5, 15)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListWithBetween, "List of products with BETWEEN condition:")
}

//Get entities list starts with LIKE conditions
func ProductsListStartsWith() {
	productsListStartsWith, err := productModel.FindProductsStartsWith(db, "Mo")

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListStartsWith, "List of products starts with LIKE condition:")
}

//Get entities list ends with LIKE conditions
func ProductsListEndsWith() {
	productsListEndsWith, err := productModel.FindProductsEndsWith(db, "vi 1")

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListEndsWith, "List of products ends with LIKE condition:")
}

//Get entities list contains with LIKE conditions
func ProductsListContains() {
	productsListContains, err := productModel.FindProductsContains(db, "bi")

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListContains, "List of products contains with LIKE condition:")
}

//Get entities list ordered by DESC
func ProductsListOrderByDesc() {
	productsListOrderByDesc, err := productModel.FindProductsOrderByDesc(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListOrderByDesc, "List of products ordered by DESC:")
}

//Get entities list ordered by ASC
func ProductsListOrderByAsc() {
	productsListOrderByAsc, err := productModel.FindProductsOrderByAsc(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListOrderByAsc, "List of products ordered by ASC:")
}

//Get entities list ordered by ASC
func ProductsListOrderByAndCondition() {
	productsListOrderByAndCondition, err := productModel.FindProductsOrderByAndCondition(db, true)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListOrderByAndCondition, "List of products ordered by DESC and with condition:")
}

//Get entities list with LIMIT
func ProductsListWithLimit() {
	productsListWithLimit, err := productModel.FindProductsWithLimit(db, 2)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListWithLimit, "List of products with LIMIT:")
}

//Get entities list ordered by with LIMIT
func ProductsListOrderByWithLimit() {
	productsListOrderByWithLimit, err := productModel.FindProductsOrderByWithLimit(db, 2)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListOrderByWithLimit, "List of products ordered by with LIMIT:")
}

//Get entities list ordered by with LIMIT and condition
func ProductsListOrderByWithWhereAndLimit() {
	productsListOrderByWithWhereAndLimit, err := productModel.FindProductsOrderByWithWhereAndLimit(db, true, 2)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListOrderByWithWhereAndLimit, "List of products ordered by with LIMIT and condition:")
}

//Get entities list by year, month and day
func ProductsListByYearAndMonthAndDay() {
	productsListByYearAndMonthAndDay, err := productModel.FindProductsByYearAndMonthAndDay(db, 2019, 11, 8)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListByYearAndMonthAndDay, "List of products by year, month and day:")
}

//Get entities list by date
func ProductsListByDate() {
	date, err := time.Parse("2006-01-02", "2019-07-10")

	if err != nil {
		log.Println(err)
		return
	}

	productsListByDate, err := productModel.FindProductsByDate(db, date)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListByDate, "List of products by date:")
}

//Get entities list by dates
func ProductsListByDates() {
	startDate, err := time.Parse("2006-01-02", "2019-07-09")

	if err != nil {
		log.Println(err)
		return
	}

	endDate, err := time.Parse("2006-01-02", "2019-07-15")

	if err != nil {
		log.Println(err)
		return
	}

	productsListByDates, err := productModel.FindProductsByDates(db, startDate, endDate)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productsListByDates, "List of products by dates:")
}

//Get entity by id
func ProductById() {
	productById, err := productModel.FindProductById(db, 2)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productById, "Product by ID:")
}

//Get entities list with select
func ProductListWithSelect() {
	productListWithSelect, err := productModel.FindProductWithSelect(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productListWithSelect, "List of products wit select:")
}

//Get entities list with select and condition
func ProductListWithSelectAndCondition() {
	productListWithSelectAndCondition, err := productModel.FindProductWithSelectAndCondition(db, true)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productListWithSelectAndCondition, "List of products with select and condition:")
}

//Get entities list with GROUP BY
func ProductListWithGroupBy() {
	productListWithGroupBy, err := productModel.FindProductGroupBy(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productListWithGroupBy, "List of products with GROUP BY:")
}

//Get entities list with HAVING
func ProductListWithHaving() {
	productListWithHaving, err := productModel.FindProductWithHaving(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(productListWithHaving, "List of products with HAVING:")
}

//Get sum of product quantities
func ProductSum() {
	result := productModel.FindSum(db)

	fmt.Println("Sum of product quantities:")
	fmt.Println("\tSum:", result)
	fmt.Println()
}

//Get sum of product quantities with condition
func ProductWithCondition() {
	result := productModel.FindSumWithCondition(db, true)

	fmt.Println("Sum of product quantities with condition:")
	fmt.Println("\tSum:", result)
	fmt.Println()
}

//Get sum of product quantities with calculate
func ProductWithCalculate() {
	result := productModel.FindSumWithCalculate(db, true)

	fmt.Println("Sum of product quantities with calculate:")
	fmt.Println("\tSum:", result)
	fmt.Println()
}

//Get count of product entities
func ProductCount() {
	result := productModel.FindCount(db)

	fmt.Println("Count of product entities:")
	fmt.Println("\tCount:", result)
	fmt.Println()
}

//Get count of product entities with condition
func ProductCountWithCondition() {
	result := productModel.FindCountWithCondition(db, true)

	fmt.Println("Count of product entities with condition:")
	fmt.Println("\tCount:", result)
	fmt.Println()
}

//Get min of product price
func ProductMin() {
	result := productModel.FindMin(db)

	fmt.Println("Min product price:")
	fmt.Println("\tPrice:", result)
	fmt.Println()
}

//Get min of product price with condition
func ProductMinWithCondition() {
	result := productModel.FindMinWithCondition(db, true)

	fmt.Println("Min product price with condition:")
	fmt.Println("\tPrice:", result)
	fmt.Println()
}

//Get max of product price
func ProductMax() {
	result := productModel.FindMax(db)

	fmt.Println("Max product price:")
	fmt.Println("\tPrice:", result)
	fmt.Println()
}

//Get max of product price with condition
func ProductMaxWithCondition() {
	result := productModel.FindMinWithCondition(db, true)

	fmt.Println("Max product price with condition:")
	fmt.Println("\tPrice:", result)
	fmt.Println()
}

//Get avg of product price
func ProductAvg() {
	result := productModel.FindAvg(db)

	fmt.Println("Avg product price:")
	fmt.Println("\tPrice:", result)
	fmt.Println()
}

//Get avg of product price with condition
func ProductAvgWithCondition() {
	result := productModel.FindAvgWithCondition(db, true)

	fmt.Println("Avg product price with condition:")
	fmt.Println("\tPrice:", result)
	fmt.Println()
}

//Create product entity
func ProductCreateEntity() {
	product := entities.Product{
		Id:       6,
		Name:     "Computer 2",
		Price:    11,
		Quantity: 3,
		Status:   true,
		Created:  time.Now(),
	}

	err := productModel.CreateEntity(db, &product)

	if err != nil {
		log.Println(err)
	} else {
		output.Output(product, "Create product entity:")
	}
}

//Update product entity
func ProductUpdateEntity() {
	product, err := productModel.FindProductByIdWithFirst(db, 6)

	if err != nil {
		log.Println(err)
	}

	product.Name = "def"
	product.Price = 999
	product.Status = false

	err = productModel.UpdateEntity(db, &product)

	if err != nil {
		log.Println(err)
	}

	output.Output(product, "Update product entity:")
}

//Delete product entity
func ProductDeleteEntity() {
	product, err := productModel.FindProductByIdWithFirst(db, 6)

	if err != nil {
		log.Println(err)
	}

	err = productModel.DeleteEntity(db, product)

	if err != nil {
		log.Println(err)
	}

	output.Output(product, "Delete product entity:")
}
