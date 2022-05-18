package controllers

import (
	"log"

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
