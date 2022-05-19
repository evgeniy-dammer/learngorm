package main

import "github.com/evgeniy-dammer/learngorm/controllers"

func main() {
	controllers.ProductsList()
	controllers.ProductsListWithCondition()
	controllers.ProductsListWithBetween()
	controllers.ProductsListStartsWith()
	controllers.ProductsListEndsWith()
	controllers.ProductsListContains()
	controllers.ProductsListOrderByDesc()
	controllers.ProductsListOrderByAsc()
	controllers.ProductsListOrderByAndCondition()
	controllers.ProductsListWithLimit()
	controllers.ProductsListOrderByWithLimit()
	controllers.ProductsListOrderByWithWhereAndLimit()
	controllers.ProductsListByYearAndMonthAndDay()
	controllers.ProductsListByDate()
	controllers.ProductsListByDates()
	controllers.ProductById()
}
