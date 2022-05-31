package entities

import "fmt"

type Role struct {
	Id           int `gorm:"primary_key, AUTO_INCREMENT"`
	Name         string
	RoleAccounts []RoleAccount `gorm:"ForeignKey:RoleID"`
}

func (role *Role) TableName() string {
	return "role"
}

func (role Role) ToString() string {
	return fmt.Sprintf("\tId: %d\tName: %s", role.Id, role.Name)
}
