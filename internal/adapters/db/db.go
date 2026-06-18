package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Adapter is a gorm-backed repository, generic over the model type T.
type Adapter[T any] struct {
	db *gorm.DB
}

// NewAdapter creates a new Adapter and auto-migrates the schema for T.
func NewAdapter[T any](connectionString string) (*Adapter[T], error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection failure: %w", err)
	}

	if err := db.AutoMigrate(new(T)); err != nil {
		return nil, fmt.Errorf("auto-migration failure: %w", err)
	}

	return &Adapter[T]{db: db}, nil
}

// Create takes a pointer of a model struct and inserts it to db
func (da Adapter[T]) Create(modelP *T) error {
	return da.db.Create(modelP).Error
}

// Find performs a query by sending a property, a value to match and a pointer to the model of the table
func (da Adapter[T]) Find(parameter string, value string, modelP *T) error {
	return da.db.Where(fmt.Sprintf("%s = ?", parameter), value).First(modelP).Error
}

// FindAll takes a pointer to a slice of models to read from db
func (da Adapter[T]) FindAll(modelsP *[]T) error {
	return da.db.Find(modelsP).Error
}
