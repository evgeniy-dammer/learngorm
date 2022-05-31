package models

import (
	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type AccountModel struct{}

//FindAllAccounts returns list of accounts
func (accountModel AccountModel) FindAllAccounts(db *gorm.DB) ([]entities.Account, error) {
	var accounts []entities.Account

	db.Preload("RoleAccounts").Preload("RoleAccounts.Role").Find(&accounts)

	return accounts, nil
}
