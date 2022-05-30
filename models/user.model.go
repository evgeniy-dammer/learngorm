package models

import (
	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type UserModel struct{}

//FindAllUsersWithLanguages returns list of users and languages relations
func (userModel UserModel) FindAllUsersWithLanguages(db *gorm.DB) ([]entities.User, error) {
	var users []entities.User

	db.Preload("Languages").Find(&users)

	return users, nil
}
