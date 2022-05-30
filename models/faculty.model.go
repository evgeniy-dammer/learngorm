package models

import (
	"github.com/evgeniy-dammer/learngorm/entities"
	"gorm.io/gorm"
)

type FacultyModel struct{}

//FindAllFacultiesWithStudents returns
func (facultyModel FacultyModel) FindAllFacultiesWithStudents(db *gorm.DB) ([]entities.Faculty, error) {
	var faculties []entities.Faculty

	db.Preload("Students").Find(&faculties)

	return faculties, nil
}

//FindAllStudentsWithFaculties returns
func (facultyModel FacultyModel) FindAllStudentsWithFaculties(db *gorm.DB) ([]entities.Student, error) {
	var students []entities.Student

	db.Preload("Faculty").Find(&students)

	return students, nil
}
