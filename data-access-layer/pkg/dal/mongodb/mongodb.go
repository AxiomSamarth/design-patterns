package mongodb

import (
	"context"

	"axiomsamarth.io/data-access-layer/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBDAL represents the MongoDB Data Access Layer.
type MongoDBDAL struct {
	collection *mongo.Collection
}

// NewMongoDBDAL creates a new instance of MongoDBDAL.
func NewMongoDBDAL(uri, dbName, collectionName string) (*MongoDBDAL, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDBDAL{collection: collection}, nil
}

// WriteStudent writes student data to MongoDB.
func (dal *MongoDBDAL) WriteStudent(student *models.Student) error {
	_, err := dal.collection.InsertOne(context.Background(), student)
	return err
}

// ReadStudents reads all student data from MongoDB.
func (dal *MongoDBDAL) ReadStudents() ([]models.Student, error) {
	var students []models.Student
	cursor, err := dal.collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err := cursor.All(context.Background(), &students); err != nil {
		return nil, err
	}
	return students, nil
}
