package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocationByZipCode_Success(t *testing.T) {
	repo := NewZipCodeRepository()

	location, err := repo.GetLocationByZipCode("12220790")
	assert.NoError(t, err)
	assert.Equal(t, "São José dos Campos", location.Localidade)
	assert.Equal(t, "Jardim Ismênia", location.Bairro)
}

func TestGetLocationByZipCode_NotFound(t *testing.T) {
	repo := NewZipCodeRepository()

	_, err := repo.GetLocationByZipCode("00000000")
	assert.Error(t, err)
	assert.Equal(t, "CEP não encontrado", err.Error())
}

func TestGetLocationByZipCode_InvalidZipCode(t *testing.T) {
	repo := NewZipCodeRepository()

	_, err := repo.GetLocationByZipCode("01153001")
	assert.Error(t, err)
	assert.Equal(t, "CEP não encontrado", err.Error())
}
