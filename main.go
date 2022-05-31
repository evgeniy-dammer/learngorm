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
	controllers.ProductListWithSelect()
	controllers.ProductListWithSelectAndCondition()
	controllers.ProductListWithGroupBy()
	controllers.ProductListWithHaving()
	controllers.ProductSum()
	controllers.ProductWithCondition()
	controllers.ProductWithCalculate()
	controllers.ProductCount()
	controllers.ProductCountWithCondition()
	controllers.ProductMin()
	controllers.ProductMinWithCondition()
	controllers.ProductMax()
	controllers.ProductMaxWithCondition()
	controllers.ProductAvg()
	controllers.ProductAvgWithCondition()
	controllers.ProductCreateEntity()
	controllers.ProductUpdateEntity()
	controllers.ProductDeleteEntity()
	controllers.FacultyListWithStudents()
	controllers.StudentsListWithFaculties()
	controllers.LanguagesListWithUsers()
	controllers.UsersListWithLanguages()
	controllers.AccountsListWithRoles()
	controllers.RolesListWithAccounts()
}
