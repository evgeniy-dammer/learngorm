package controllers

import (
	"log"

	"github.com/evgeniy-dammer/learngorm/models"
	"github.com/evgeniy-dammer/learngorm/output"
)

var roleModel models.RoleModel

//Get entities list using composition of primary keys
func RolesListWithAccounts() {
	rolessList, err := roleModel.FindAllRoles(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(rolessList, "List of all roles:")
}
