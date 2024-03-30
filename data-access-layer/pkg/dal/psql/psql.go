package psql

import (
	"axiomsamarth.io/data-access-layer/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDAL represents the PostgreSQL Data Access Layer.
type PostgresDAL struct {
	db *gorm.DB
}

// NewPostgresDAL creates a new instance of PostgresDAL.
func NewPostgresDAL(dsn string) (*PostgresDAL, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDAL{db: db}, nil
}

// WriteStudent writes student data to PostgreSQL.
func (dal *PostgresDAL) WriteStudent(student *models.Student) error {
	return dal.db.Create(student).Error
}

// ReadStudents reads all student data from PostgreSQL.
func (dal *PostgresDAL) ReadStudents() ([]models.Student, error) {
	var students []models.Student
	if err := dal.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
