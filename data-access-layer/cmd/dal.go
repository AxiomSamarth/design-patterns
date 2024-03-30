package main

import (
	"flag"
	"fmt"

	"axiomsamarth.io/data-access-layer/cmd/app"
	"axiomsamarth.io/data-access-layer/pkg/dal"
	"axiomsamarth.io/data-access-layer/pkg/models"

	"github.com/pkg/errors"
)

func main() {
	var dbProvider string

	flag.StringVar(&dbProvider, "db-provider", "psql", "Specify the name of the DB Provider to use ['psql', 'mongodb']")
	flag.Parse()

	dalDriver, err := app.GetDALDriver(dbProvider)
	if err != nil {
		fmt.Printf("error initializing DAL Driver: %v", err)
		return
	}

	collegeControl := dal.NewCollegeControl(dalDriver)
	err = useCollegeControl(collegeControl)
	if err != nil {
		fmt.Printf("Error using College Control with PostgreSQL DAL: %v\n", err)
		return
	}
}

func useCollegeControl(cc *dal.CollegeControl) error {
	// Add a student
	student := &models.Student{
		Name:  "Alice",
		Age:   20,
		Grade: 85,
	}
	if err := cc.AddStudent(student); err != nil {
		return errors.Wrap(err, "error adding student")
	}

	// List all students
	students, err := cc.ListStudents()
	if err != nil {
		return errors.Wrap(err, "error listing students")
	}
	fmt.Println("Students:")
	for _, s := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %d\n", s.ID, s.Name, s.Age, s.Grade)
	}
	return nil
}
