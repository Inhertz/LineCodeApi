package application

import (
	"LineCodeApi/internal/core/domain"
	"LineCodeApi/internal/core/models"
)

// Application implements the APIPort interface
type Application struct {
	logic *domain.Logic
	db    DbPort
}

// NewApplication creates a new Application
func NewApplication(db DbPort, logic *domain.Logic) *Application {
	return &Application{db: db, logic: logic}
}

// GetAllManchester is an use case where all manchester codes stored are fetched
func (apia Application) GetAllManchester() ([]models.Manchester, error) {
	var manchesters []models.Manchester
	err := apia.db.FindAll(&manchesters)
	if err != nil {
		return nil, err
	}
	return manchesters, nil
}

// GenerateEncodedManchester is an use case where a manchester encoding is performed and saved
func (apia Application) GenerateEncodedManchester(manchester *models.Manchester) error {
	err := apia.db.Find("decoded", manchester.Decoded, manchester)
	if err == nil {
		return nil
	}

	encoded, err := apia.logic.ManchesterEncode(manchester.Decoded)
	if err != nil {
		return err
	}

	manchester.Encoded = encoded
	manchester.EncodedPulseWidth = manchester.DecodedPulseWidth / 2

	err = apia.db.Create(manchester)
	if err != nil {
		return err
	}
	return nil
}

// GenerateDecodedManchester is an use case where a manchester encoding is performed and saved
func (apia Application) GenerateDecodedManchester(manchester *models.Manchester) error {
	decoded, err := apia.logic.ManchesterDecode(manchester.Encoded)
	if err != nil {
		return err
	}

	manchester.Decoded = decoded
	manchester.DecodedPulseWidth = manchester.EncodedPulseWidth * 2

	err = apia.db.Create(manchester)
	if err != nil {
		return err
	}
	return nil
}
