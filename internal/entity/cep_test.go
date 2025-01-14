package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCEP(t *testing.T) {
	cep := NewCEP("12220780", "Rua Bapendi", "", "Jardim Ismênia", "São José dos Campos", "SP", "3549904", "6452", "12", "7099")

	assert.Equal(t, cep.CEP, "12220780")
	assert.Equal(t, cep.Logradouro, "Rua Bapendi")
	assert.Equal(t, cep.Complemento, "")
	assert.Equal(t, cep.Bairro, "Jardim Ismênia")
	assert.Equal(t, cep.Localidade, "São José dos Campos")
	assert.Equal(t, cep.UF, "SP")
	assert.Equal(t, cep.IBGE, "3549904")
	assert.Equal(t, cep.GIA, "6452")
	assert.Equal(t, cep.DDD, "12")
	assert.Equal(t, cep.SIAFI, "7099")
}
