package mocks

import (
	"LineCodeApi/internal/core/models"
	"errors"
)

type ApplicationMock struct {
	ExpectedOutput interface{}
}

// GetAllManchester is an use case where all manchester codes stored are fetched
func (a ApplicationMock) GetAllManchester() ([]models.Manchester, error) {
	manchesters, ok := a.ExpectedOutput.([]models.Manchester)
	if !ok {
		return nil, errors.New("bad mock input")
	}
	return manchesters, nil
}

// GenerateEncodedManchester is an use case where a manchester encoding is performed and saved
func (a ApplicationMock) GenerateEncodedManchester(manchester *models.Manchester) error {
	expected, ok := a.ExpectedOutput.(models.Manchester)
	if !ok {
		return errors.New("bad mock input")
	}
	manchester.ID = expected.ID
	manchester.Decoded = expected.Decoded
	manchester.Encoded = expected.Encoded
	manchester.DecodedPulseWidth = expected.DecodedPulseWidth
	manchester.EncodedPulseWidth = expected.EncodedPulseWidth
	manchester.Unit = expected.Unit

	return nil
}

// GenerateDecodedManchester is an use case where a manchester encoding is performed and saved
func (a ApplicationMock) GenerateDecodedManchester(manchester *models.Manchester) error {
	expected, ok := a.ExpectedOutput.(models.Manchester)
	if !ok {
		return errors.New("bad mock input")
	}
	manchester.ID = expected.ID
	manchester.Decoded = expected.Decoded
	manchester.Encoded = expected.Encoded
	manchester.DecodedPulseWidth = expected.DecodedPulseWidth
	manchester.EncodedPulseWidth = expected.EncodedPulseWidth
	manchester.Unit = expected.Unit

	return nil
}
