package controllers

import (
	"log"

	"github.com/evgeniy-dammer/learngorm/models"
	"github.com/evgeniy-dammer/learngorm/output"
)

var userModel models.UserModel

//Get entities list using many to many relationship
func UsersListWithLanguages() {
	usersList, err := userModel.FindAllUsersWithLanguages(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(usersList, "List of all languages:")
}
