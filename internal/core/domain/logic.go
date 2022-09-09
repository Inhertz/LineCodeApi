package domain

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"
)

// Logic exports the domain logic to the application layer
type Logic struct {
	ManchesterEncodeDictionary map[byte]string
	ManchesterDecodeDictionary map[string]byte
}

// New creates a new Logic struct
func New() *Logic {

	manEncDic, manDecDic := newManchesterDictionaries()

	return &Logic{
		ManchesterEncodeDictionary: manEncDic,
		ManchesterDecodeDictionary: manDecDic,
	}
}

// newManchesterDictionary creates a map which holds the Manchester encoding of all 256 bytes
func newManchesterDictionaries() (map[byte]string, map[string]byte) {
	encoded := make(map[byte]string)
	decoded := make(map[string]byte)
	for i := 0; i < 256; i++ {
		byteIn := uint8(i)
		pulses := byteToManchester(byteIn)
		encoded[uint8(i)] = pulses
		decoded[pulses] = byteIn
	}

	return encoded, decoded
}

// byteToManchester generates an int slice from a single byte or uint8
func byteToManchester(in byte) string {
	var out string
	for i := 0; i < 8; i++ {
		mask := in & uint8(math.Pow(2, float64(7-i)))
		if mask == 0 {
			out = out + "-A" + "+A"
		} else {
			out = out + "+A" + "-A"
		}
	}
	return out
}

// ManchesterEncode takes a string representation of a byte array in Hex and returns its Manchester code
func (l Logic) ManchesterEncode(in string) (string, error) {
	var out string
	arr, err := hex.DecodeString(in)
	if err != nil {
		return "", err
	}

	for _, value := range arr {
		out = out + l.ManchesterEncodeDictionary[value]
	}

	return out, nil
}

// ManchesterDecode takes a string representation of a Manchester code and returns its byte array in Hex
func (l Logic) ManchesterDecode(in string) (string, error) {
	var out string
	err := checkManchesterCode(in)
	if err != nil {
		return "", err
	}

	for i := 0; i < len([]rune(in)); i += 32 {
		pulses := in[i : i+32]
		byteOut := l.ManchesterDecodeDictionary[pulses]
		if byteOut == 0 && pulses != "-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A" {
			return "", fmt.Errorf("wrong pulse inside byte: %d", i)
		}
		out = out + hex.EncodeToString([]byte{byteOut})
	}

	return out, nil
}

// checkManchesterCode returns an error if the number of pulses is not multiple of 8
func checkManchesterCode(in string) error {
	totalPulses := strings.Count(in, "A")
	if totalPulses%8 != 0 {
		return errors.New("number of pulses should be a multiple of 8 in order to return bytes")
	}

	return nil
}
