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
	case []entities.Faculty:
		for _, faculty := range ent {
			fmt.Println(faculty.ToString())
			if len(faculty.Students) > 0 {
				for _, student := range faculty.Students {
					fmt.Println("\t" + student.ToString())
				}
			}
		}
	case []entities.Student:
		for _, student := range ent {
			fmt.Println(student.ToString())
			fmt.Println("\t" + student.Faculty.ToString())
		}
	case []entities.Language:
		for _, language := range ent {
			fmt.Println(language.ToString())
			if len(language.Users) > 0 {
				for _, user := range language.Users {
					fmt.Println("\t" + user.ToString())
				}
			}
		}
	case []entities.User:
		for _, user := range ent {
			fmt.Println(user.ToString())
			if len(user.Languages) > 0 {
				for _, language := range user.Languages {
					fmt.Println("\t" + language.ToString())
				}
			}
		}
	case []entities.Account:
		for _, account := range ent {
			fmt.Println(account.ToString())
			if len(account.RoleAccounts) > 0 {
				for _, role := range account.RoleAccounts {
					fmt.Println("\t" + role.ToString())
				}
			}
		}
	case []entities.Role:
		for _, role := range ent {
			fmt.Println(role.ToString())
			if len(role.RoleAccounts) > 0 {
				for _, account := range role.RoleAccounts {
					fmt.Println("\t" + account.ToString())
				}
			}
		}
	}

	fmt.Println()
}
