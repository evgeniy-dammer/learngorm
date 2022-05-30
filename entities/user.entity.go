package entities

import "fmt"

type User struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Username  string
	Password  string
	FullName  string     `gorm:"column:full_name"`
	Languages []Language `gorm:"many2many:user_language"`
}

func (user *User) TableName() string {
	return "user"
}

func (user User) ToString() string {
	return fmt.Sprintf("\tId: %d\tUsername: %s\tFull Name: %s", user.Id, user.Username, user.FullName)
}
