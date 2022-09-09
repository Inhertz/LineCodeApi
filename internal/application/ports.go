package application

import "LineCodeApi/internal/core/models"

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	GetAllManchester() ([]models.Manchester, error)
	GenerateEncodedManchester(manchester *models.Manchester) error
	GenerateDecodedManchester(manchester *models.Manchester) error
}

// DbPort is the port for a db adapter
type DbPort interface {
	Create(modelP interface{}) error
	Find(parameter string, value string, modelP interface{}) error
	FindAll(modelsP interface{}) error
}

// GRPCPort interface is implemented with protoc inside the adapter!!!
