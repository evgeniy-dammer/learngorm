package main

import (
	"fmt"
	"log"

	"github.com/evgeniy-dammer/learngorm/config"
	"github.com/evgeniy-dammer/learngorm/models"
	"gorm.io/gorm"
)

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

func main() {
	var productModel models.ProductModel

	//Get entities list from Database
	productsList, err := productModel.FindAll(db)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nList of all products:")

	for _, product := range productsList {
		fmt.Println(product.ToString())
		fmt.Println()
	}

	//Get entities list with conditions from Database
	productsListWithCondition, err := productModel.FindProductsWithConditions(db, true, 5, 15)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nList of products with condition:")

	for _, product := range productsListWithCondition {
		fmt.Println(product.ToString())
		fmt.Println()
	}

	//Get entities list with BETWEEN conditions from Database
	productsListWithBetween, err := productModel.FindProductsWithBetween(db, 5, 15)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nList of products with BETWEEN condition:")

	for _, product := range productsListWithBetween {
		fmt.Println(product.ToString())
		fmt.Println()
	}
}
