package entities

import "fmt"

type Faculty struct {
	Id       int `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string
	Students []Student `gorm:"ForeignKey:FacultyID"`
}

func (faculty *Faculty) TableName() string {
	return "faculty"
}

func (faculty Faculty) ToString() string {
	return fmt.Sprintf("\tId: %d\tName: %s", faculty.Id, faculty.Name)
}
