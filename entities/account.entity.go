package entities

import "fmt"

type Account struct {
	Id           int `gorm:"primary_key, AUTO_INCREMENT"`
	Username     string
	Password     string
	RoleAccounts []RoleAccount `gorm:"ForeignKey:AccountID"`
}

func (account *Account) TableName() string {
	return "account"
}

func (account Account) ToString() string {
	return fmt.Sprintf("\tId: %d\tUsername: %s", account.Id, account.Username)
}
