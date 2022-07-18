package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryption(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		except error
	}{
		{"case 1", "12345678", nil},
		{"case 2", "1234lmh", nil},
		{"case 3", "123lmh...", nil},
		{"case 4", "1234pkp...$$$###", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Encryption(tt.param)
			want := tt.except

			assert.Equal(t, want, got)
		})
	}
}

func TestInterpretation(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		except bool
	}{
		{"case 1", "12345678", "", true},
		{"case 2", "1234lmh", "", true},
		{"case 3", "123lmh...", "", true},
		{"case 4", "1234pkp...$$$###", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, tt.param2 = Encryption(tt.param1)
			want := tt.except

			err, got := Interpretation(tt.param2, tt.param1)

			assert.Nil(t, err)
			assert.Equal(t, want, got)
		})
	}
}
