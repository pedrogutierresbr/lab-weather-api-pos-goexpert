package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	repo := NewCEPRepository()

	assert.True(t, repo.IsValidCEP("38050600"))
	assert.False(t, repo.IsValidCEP("380506"))
	assert.False(t, repo.IsValidCEP("380506000"))
}

func TestGetCEP(t *testing.T) {
	repository := NewCEPRepository()
	cep_data, err := repository.GetCEP("38050600")
	assert.Nil(t, err)
	assert.NotNil(t, cep_data)
	assert.Contains(t, string(cep_data), "Uberaba")
}

func TestConvertResponse(t *testing.T) {
	response := []byte(`{
		"cep": "38050-600",
		"logradouro": "Rua Ronan Martins Marques",
		"complemento": "",
		"unidade": "",
		"bairro": "Santa Maria",
		"localidade": "Uberaba",
		"uf": "MG",
		"estado": "Minas Gerais",
		"regiao": "Sudeste",
		"ibge": "3170107",
		"gia": "",
		"ddd": "34",
		"siafi": "5401"
	}`)

	repository := NewCEPRepository()
	cep, err := repository.ConvertResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, cep.CEP, "38050-600")
	assert.Equal(t, cep.Logradouro, "Rua Ronan Martins Marques")
	assert.Equal(t, cep.Bairro, "Santa Maria")
	assert.Equal(t, cep.Localidade, "Uberaba")
	assert.Equal(t, cep.UF, "MG")
}
