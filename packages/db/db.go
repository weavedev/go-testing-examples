package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq" // Postgres driver

	"github.com/jinzhu/gorm"
)

const (
	dbConnection = "dbname=%s user=%s password=%s port=%d host=%s sslmode=disable"
)

// ConnectTest illustratioin of connecting to a db
func ConnectTest(models []interface{}) (*gorm.DB, error) {
	return Connect(models, "presentation_test", os.Getenv("PG_USERNAME"), os.Getenv("PG_PASSWORD"), 5432, "localhost")
}

// Connect creates connection to postgres database
func Connect(models []interface{}, dbName, dbUser, dbPass string, dbPort int, dbHost string) (*gorm.DB, error) {
	url := fmt.Sprintf(dbConnection, dbName, dbUser, dbPass, dbPort, dbHost)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(models...)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "pg_trgm";`)

	return db, nil
}
