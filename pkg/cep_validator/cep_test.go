package cepvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCEP(t *testing.T) {
	tests := []struct {
		name     string
		inputCEP string
		expected bool
	}{
		{"CEP válido com hífen", "12345-678", true},
		{"CEP válido sem hífen", "12345678", true},
		{"CEP inválido com letras", "12345-ABC", false},
		{"CEP inválido com menos dígitos", "1234-567", false},
		{"CEP inválido com mais dígitos", "123456789", false},
		{"CEP vazio", "", false},
		{"CEP com caracteres especiais", "12-345-678", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidCEP(tt.inputCEP)
			assert.Equal(t, tt.expected, result, "Resultado para %s deveria ser %v", tt.inputCEP, tt.expected)
		})
	}
}
