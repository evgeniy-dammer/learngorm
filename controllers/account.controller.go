package controllers

import (
	"log"

	"github.com/evgeniy-dammer/learngorm/models"
	"github.com/evgeniy-dammer/learngorm/output"
)

var accountModel models.AccountModel

//Get entities list using composition of primary keys
func AccountsListWithRoles() {
	accountList, err := accountModel.FindAllAccounts(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(accountList, "List of all accounts:")
}
