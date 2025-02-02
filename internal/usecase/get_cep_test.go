package usecase

import (
	"testing"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCEPRepository struct {
	mock.Mock
}

func (m *MockCEPRepository) IsValidCEP(cep string) bool {
	args := m.Called(cep)
	return args.Bool(0)
}

func (m *MockCEPRepository) ConvertResponse(cepData []byte) (*entity.CEP, error) {
	args := m.Called(cepData)
	return args.Get(0).(*entity.CEP), args.Error(1)
}

func (m *MockCEPRepository) GetCEP(cep string) ([]byte, error) {
	args := m.Called(cep)
	return args.Get(0).([]byte), args.Error(1)
}

func TestGetCEPUsecase(t *testing.T) {
	mockCEPRepo := new(MockCEPRepository)

	mockCEPRepo.On("IsValidCEP", "38050600").Return(true)
	mockCEPRepo.On("IsValidCEP", "00000000").Return(false)

	mockCEPRepo.On("GetCEP", "38050600").Return([]byte(`{
		"cep": "38050-600",
		"logradouro": "Rua Ronan Martins Marques",
		"complemento": "",
		"bairro": "Santa Maria",
		"localidade": "Uberaba",
		"uf": "MG",
		"ibge": "3170107",
		"gia": "",
		"ddd": "34",
		"siafi": "5401"
	}`), nil)
	mockCEPRepo.On("ConvertResponse", mock.Anything).Return(&entity.CEP{
		CEP:         "38050600",
		Logradouro:  "Rua Ronan Martins Marques",
		Complemento: "",
		Bairro:      "Santa Maria",
		Localidade:  "Uberaba",
		UF:          "MG",
		IBGE:        "3170107",
		GIA:         "",
		DDD:         "34",
		SIAFI:       "5401",
	}, nil)

	usecase := NewCepUseCase(mockCEPRepo)

	input := CEPInputDTO{CEP: "38050600"}
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, "38050600", output.CEP)
	assert.Equal(t, "Rua Ronan Martins Marques", output.Logradouro)
	assert.Equal(t, "Santa Maria", output.Bairro)
	assert.Equal(t, "Uberaba", output.Localidade)

	input = CEPInputDTO{CEP: "00000000"}
	output, err = usecase.Execute(input)
	assert.NotNil(t, err)
	assert.Equal(t, "CEP inválido", err.Error())
	assert.Empty(t, output)

	mockCEPRepo.AssertExpectations(t)
}
