package entities

import "fmt"

type RoleAccount struct {
	Id        int
	AccountID int `gorm:"primary_key; column:account_id"`
	RoleID    int `gorm:"primary_key; column:role_id"`
	Status    bool
	Role      Role
	Account   Account
}

func (roleAccount *RoleAccount) TableName() string {
	return "role_account"
}

func (roleAccount RoleAccount) ToString() string {
	return fmt.Sprintf("\tRole Id: %d\tAccount Id: %d\tStatus: %t", roleAccount.RoleID, roleAccount.AccountID, roleAccount.Status)
}
