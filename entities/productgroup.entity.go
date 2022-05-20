package entities

import "fmt"

type ProductGroup struct {
	Status  bool
	Result1 int
	Result2 int
	Result3 float64
	Result4 float64
	Result5 float64
}

func (productGroup ProductGroup) ToString() string {
	return fmt.Sprintf(
		"\tStatus: %t\tCountProducts: %d\tSumQuantities: %d\tMinPrice: %0.1f\tMaxPrice: %0.1f\tAvgPrice: %0.1f",
		productGroup.Status,
		productGroup.Result1,
		productGroup.Result2,
		productGroup.Result3,
		productGroup.Result4,
		productGroup.Result5,
	)
}
