package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByteToManchester(t *testing.T) {
	byteToManchesterPairs := []struct {
		name     string
		input    byte
		expected string
	}{
		{"OK 0b0000 0000", 0, "-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A"},
		{"OK 0b1111 1111", 255, "+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A"},
		{"OK 217", 217, "+A-A+A-A-A+A+A-A+A-A-A+A-A+A+A-A"},
	}

	for _, pair := range byteToManchesterPairs {
		ans := byteToManchester(pair.input)
		require.Equal(t, pair.expected, ans)
	}
}

func TestManchesterEncode(t *testing.T) {
	manchesterEncodePairs := []struct {
		name     string
		input    string
		expected string
	}{
		{"OK 0b0000 0000", "00", "-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A"},
		{"OK 0xff", "ff", "+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A"},
		{"OK 0xFF", "FF", "+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A"},
		{"OK 0xFF00", "FF00", "+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A"},
		{"OK 0xAE01", "AE01", "+A-A-A+A+A-A-A+A+A-A+A-A+A-A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A+A-A"},
	}

	dom := New()

	for _, pair := range manchesterEncodePairs {
		ans, err := dom.ManchesterEncode(pair.input)
		if err != nil {
			t.Errorf("expected: %v, got %v", nil, err)
		}
		require.Equal(t, pair.expected, ans, pair.name)
	}
}

func TestManchesterDecode(t *testing.T) {
	manchesterDecodePairs := []struct {
		name     string
		input    string
		expected string
	}{
		{"OK 0b0000 0000", "-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A", "00"},
		{"OK 0xff", "+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A", "ff"},
		{"OK 0xff00", "+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A", "ff00"},
		{"OK 0xae01", "+A-A-A+A+A-A-A+A+A-A+A-A+A-A-A+A-A+A-A+A-A+A-A+A-A+A-A+A-A+A+A-A", "ae01"},
	}

	dom := New()

	for _, pair := range manchesterDecodePairs {
		ans, err := dom.ManchesterDecode(pair.input)
		if err != nil {
			t.Errorf("expected: %v, got %v", nil, err)
		}
		require.Equal(t, pair.expected, ans, pair.name)
	}
}
