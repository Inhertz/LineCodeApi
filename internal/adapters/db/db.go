package db

import (
	"LineCodeApi/internal/core/models"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

// NewAdapter creates a new Adapter
func NewAdapter(connectionString string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	db.AutoMigrate(&models.Manchester{})
	return &Adapter{db: db}, nil
}

// Create takes a pointer of a model struct and inserts it to db
func (da Adapter) Create(modelP interface{}) error {
	err := da.db.Create(modelP).Error
	if err != nil {
		return err
	}

	return nil
}

// Find performs a query by sending a property, a value to match and a pointer to the model of the table
func (da Adapter) Find(parameter string, value string, modelP interface{}) error {
	err := da.db.Where(fmt.Sprintf("%s = ?", parameter), value).First(modelP).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll takes a pointer to an array of models to read from db
func (da Adapter) FindAll(modelsP interface{}) error {
	err := da.db.Find(modelsP).Error
	if err != nil {
		return err
	}

	return nil
}
