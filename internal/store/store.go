package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB opens a GORM DB connection using a Postgres DSN.
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}