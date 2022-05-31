package models

import (
	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type RoleModel struct{}

//FindAllRoles returns list of roles
func (roleModel RoleModel) FindAllRoles(db *gorm.DB) ([]entities.Role, error) {
	var roles []entities.Role

	db.Preload("RoleAccounts").Preload("RoleAccounts.Account").Find(&roles)

	return roles, nil
}
