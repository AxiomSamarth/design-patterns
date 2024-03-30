package dal

import "axiomsamarth.io/data-access-layer/pkg/models"

// DataAccessLayer represents the Data Access Layer interface.
type DataAccessLayer interface {
	WriteStudent(student *models.Student) error
	ReadStudents() ([]models.Student, error)
}

// CollegeControl represents the college control layer.
type CollegeControl struct {
	dal DataAccessLayer
}

// NewCollegeControl creates a new instance of CollegeControl.
func NewCollegeControl(dal DataAccessLayer) *CollegeControl {
	return &CollegeControl{dal: dal}
}

// AddStudent adds a new student to the college database.
func (cc *CollegeControl) AddStudent(student *models.Student) error {
	return cc.dal.WriteStudent(student)
}

// ListStudents lists all students in the college database.
func (cc *CollegeControl) ListStudents() ([]models.Student, error) {
	return cc.dal.ReadStudents()
}
