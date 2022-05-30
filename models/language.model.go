package models

import (
	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type LanguageModel struct{}

//FindAllLanguagesWithUsers returns list of languages and user relations
func (languageModel LanguageModel) FindAllLanguagesWithUsers(db *gorm.DB) ([]entities.Language, error) {
	var languages []entities.Language

	db.Preload("Users").Find(&languages)

	return languages, nil
}
