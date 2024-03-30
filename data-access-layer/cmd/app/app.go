package app

import (
	"github.com/pkg/errors"

	"axiomsamarth.io/data-access-layer/pkg/dal"
	"axiomsamarth.io/data-access-layer/pkg/dal/mongodb"
	"axiomsamarth.io/data-access-layer/pkg/dal/psql"
)

// GetDALDriver will initialize the DAL Driver of the respective DB provider.
func GetDALDriver(dbProvider string) (dal.DataAccessLayer, error) {
	switch dbProvider {
	case "psql":
		postgresDAL, err := psql.NewPostgresDAL("postgres://admin:password@localhost:5432/college?sslmode=disable")
		if err != nil {
			return nil, errors.Wrap(err, "error creating psql driver")
		}
		return postgresDAL, nil
	case "mongodb":
		mongoDAL, err := mongodb.NewMongoDBDAL("mongodb://admin:password@localhost:27017", "college", "students")
		if err != nil {
			return nil, errors.Wrap(err, "error creating mongodb driver")
		}
		return mongoDAL, nil
	default:
		return nil, errors.New("unknown database provider")
	}
}
