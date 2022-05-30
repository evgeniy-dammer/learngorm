package controllers

import (
	"log"

	"github.com/evgeniy-dammer/learngorm/models"
	"github.com/evgeniy-dammer/learngorm/output"
)

var languageModel models.LanguageModel

//Get entities list using many to many relationship
func LanguagesListWithUsers() {
	languagesList, err := languageModel.FindAllLanguagesWithUsers(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(languagesList, "List of all languages:")
}
