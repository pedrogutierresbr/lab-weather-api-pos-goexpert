package usecase

import (
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
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
