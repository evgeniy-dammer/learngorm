package controllers

import (
	"log"

	"github.com/evgeniy-dammer/learngorm/models"
	"github.com/evgeniy-dammer/learngorm/output"
)

var facultyModel models.FacultyModel

//Get entities list using one to many relationship
func FacultyListWithStudents() {
	facultyList, err := facultyModel.FindAllFacultiesWithStudents(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(facultyList, "List of all faculties:")
}

//Get entities list using many to one relationship
func StudentsListWithFaculties() {
	studentList, err := facultyModel.FindAllStudentsWithFaculties(db)

	if err != nil {
		log.Println(err)
		return
	}

	output.Output(studentList, "List of all students:")
}
