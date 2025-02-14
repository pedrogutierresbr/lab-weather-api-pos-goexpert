package usecase

import (
	"errors"
	"testing"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) GetWeather(location string) (map[string]float64, error) {
	args := m.Called(location)
	return args.Get(0).(map[string]float64), args.Error(1)
}

func TestWeatherUseCase_GetWeatherByZipCode(t *testing.T) {
	zipCodeRepo := &repository.MockZipCodeRepository{}
	weatherService := &MockWeatherService{}
	useCase := NewWeatherUseCase(zipCodeRepo, weatherService)

	t.Run("success", func(t *testing.T) {
		zipCodeRepo.On("GetLocationByZipCode", "38050600").Return(&repository.Location{
			Localidade: "Uberaba",
			Bairro:     "Santa Maria",
		}, nil)
		weatherService.On("GetWeather", "Santa Maria, Uberaba").Return(map[string]float64{
			"temp_C": 25.0,
			"temp_F": 77.0,
			"temp_K": 298.15,
		}, nil)

		weather, err := useCase.GetWeatherByZipCode("38050600")

		assert.NoError(t, err)
		assert.Equal(t, 25.0, weather["temp_C"])
		assert.Equal(t, 77.0, weather["temp_F"])
		assert.Equal(t, 298.15, weather["temp_K"])

		zipCodeRepo.AssertExpectations(t)
		weatherService.AssertExpectations(t)
	})

	t.Run("zip code not found", func(t *testing.T) {
		zipCodeRepo.On("GetLocationByZipCode", "00000000").Return(nil, errors.New("CEP não encontrado"))

		_, err := useCase.GetWeatherByZipCode("00000000")

		assert.Error(t, err)
		assert.Equal(t, "CEP não encontrado", err.Error())

		zipCodeRepo.AssertExpectations(t)
	})

	t.Run("invalid zip code", func(t *testing.T) {
		zipCodeRepo.On("GetLocationByZipCode", "invalid").Return(nil, errors.New("CEP não encontrado"))

		_, err := useCase.GetWeatherByZipCode("invalid")

		assert.Error(t, err)
		assert.Equal(t, "CEP não encontrado", err.Error())

		zipCodeRepo.AssertExpectations(t)
	})
}
