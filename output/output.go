package output

import (
	"fmt"

	"github.com/evgeniy-dammer/learngorm/entities"
)

//Output print list of entities in console
func Output(input interface{}, listName string) {
	fmt.Println(listName)

	switch ent := input.(type) {
	case []entities.ProductGroup:
		for _, productGroup := range ent {
			fmt.Println(productGroup.ToString())
		}
	case []entities.Product:
		for _, product := range ent {
			fmt.Println(product.ToString())
		}
	case entities.Product:
		fmt.Println(ent.ToString())
	}

	fmt.Println()
}
