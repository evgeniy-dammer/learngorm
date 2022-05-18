package output

import (
	"fmt"

	"github.com/evgeniy-dammer/learngorm/entities"
)

//Output print list of products in console
func Output(products []entities.Product, listName string) {
	fmt.Println(listName)

	for _, product := range products {
		fmt.Println(product.ToString())
	}

	fmt.Println()
}
