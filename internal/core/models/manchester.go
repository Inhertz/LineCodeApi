package models

type Manchester struct {
	ID                int     `json:"id"`
	Decoded           string  `json:"decoded" gorm:"unique"`
	Encoded           string  `json:"encoded" gorm:"unique"`
	DecodedPulseWidth float64 `json:"decodedPulseWidth"`
	EncodedPulseWidth float64 `json:"encodedPulseWidth"`
	Unit              string  `json:"unit"`
}
