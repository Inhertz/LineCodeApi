package application

import "LineCodeApi/internal/core/models"

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	GetAllManchester() ([]models.Manchester, error)
	GenerateEncodedManchester(manchester *models.Manchester) error
	GenerateDecodedManchester(manchester *models.Manchester) error
}

// DbPort is the port for a db adapter, generic over the model type T.
// Method-level type parameters are not allowed on interfaces in Go, so the
// type parameter lives on the interface itself.
type DbPort[T any] interface {
	Create(modelP *T) error
	Find(parameter string, value string, modelP *T) error
	FindAll(modelsP *[]T) error
}

// GRPCPort interface is implemented with protoc inside the adapter!!!
