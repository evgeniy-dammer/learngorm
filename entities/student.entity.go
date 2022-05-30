package entities

import "fmt"

type Student struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Name      string
	Address   string
	FacultyID int `gorm:"column:faculty_id"`
	Faculty   Faculty
}

func (student *Student) TableName() string {
	return "students"
}

func (student Student) ToString() string {
	return fmt.Sprintf("\tId: %d\tName: %s\tAddress: %s", student.Id, student.Name, student.Address)
}
