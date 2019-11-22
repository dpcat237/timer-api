package db

import (
	bolt "go.etcd.io/bbolt"

	"gitlab.com/dpcat237/timer-api/model"
)

const dbPath = "repository/"

// DatabaseCollector defines required methods for collector of database connections
type DatabaseCollector interface {
	GetDatabase() *bolt.DB
	GracefulStop()
}

// databaseCollector defines database connections
type databaseCollector struct {
	Db *bolt.DB
}

// GetDatabase returns database connection
func (cll *databaseCollector) GetDatabase() *bolt.DB {
	return cll.Db
}

// GracefulStop stops database connections
func (cll *databaseCollector) GracefulStop() {
	cll.Db.Close()
}

// InitDbCollector initializes database connections and set to collector
func InitDbCollector(dbName string) (DatabaseCollector, model.Error) {
	var cll databaseCollector
	path := dbPath + dbName
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return &cll, model.NewErrorServer("Error opening database file").WithError(err)
	}
	cll.Db = db
	return &cll, model.NewErrorNil()
}
