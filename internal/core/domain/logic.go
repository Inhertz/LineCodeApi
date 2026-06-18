package domain

import (
	"encoding/hex"
	"errors"
	"fmt"
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

// byteToManchester generates the Manchester pulse string for a single byte
func byteToManchester(in byte) string {
	var b strings.Builder
	b.Grow(32)
	for i := 0; i < 8; i++ {
		mask := in & (1 << (7 - i))
		if mask == 0 {
			b.WriteString("-A+A")
		} else {
			b.WriteString("+A-A")
		}
	}
	return b.String()
}

// ManchesterEncode takes a string representation of a byte array in Hex and returns its Manchester code
func (l Logic) ManchesterEncode(in string) (string, error) {
	arr, err := hex.DecodeString(in)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.Grow(len(arr) * 32)
	for _, value := range arr {
		b.WriteString(l.ManchesterEncodeDictionary[value])
	}

	return b.String(), nil
}

// ManchesterDecode takes a string representation of a Manchester code and returns its byte array in Hex
func (l Logic) ManchesterDecode(in string) (string, error) {
	if err := checkManchesterCode(in); err != nil {
		return "", err
	}

	var b strings.Builder
	b.Grow(len(in) / 16)
	for i := 0; i+32 <= len(in); i += 32 {
		pulses := in[i : i+32]
		byteOut, ok := l.ManchesterDecodeDictionary[pulses]
		if !ok {
			return "", fmt.Errorf("wrong pulse inside byte: %d", i)
		}
		b.WriteString(hex.EncodeToString([]byte{byteOut}))
	}

	return b.String(), nil
}

// checkManchesterCode returns an error if the number of pulses is not multiple of 8
func checkManchesterCode(in string) error {
	totalPulses := strings.Count(in, "A")
	if totalPulses%8 != 0 {
		return errors.New("number of pulses should be a multiple of 8 in order to return bytes")
	}

	return nil
}
