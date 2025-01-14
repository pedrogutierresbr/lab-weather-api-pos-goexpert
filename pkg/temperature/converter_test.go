package temperature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name         string
		celsius      float64
		expectedFahr float64
	}{
		{"Zero Celsius", 0, 32},
		{"Temperatura negativa", -40, -40},
		{"Temperatura positiva", 100, 212},
		{"Temperatura decimal", 36.6, 97.88},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertCelsiusToFahrenheit(tt.celsius)
			assert.InDelta(t, tt.expectedFahr, result, 0.01, "Resultado para %f deveria ser %f", tt.celsius, tt.expectedFahr)
		})
	}
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name        string
		celsius     float64
		expectedKel float64
	}{
		{"Zero Celsius", 0, 273},
		{"Temperatura negativa", -273.15, -0.15},
		{"Temperatura positiva", 100, 373},
		{"Temperatura decimal", 25.5, 298.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertCelsiusToKelvin(tt.celsius)
			assert.InDelta(t, tt.expectedKel, result, 0.01, "Resultado para %f deveria ser %f", tt.celsius, tt.expectedKel)
		})
	}
}
