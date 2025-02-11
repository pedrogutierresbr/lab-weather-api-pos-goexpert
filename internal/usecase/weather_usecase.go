package usecase

import (
	"errors"
	"fmt"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/repository"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/services"
)

type WeatherUseCase struct {
	zipCodeRepo    repository.ZipCodeRepository
	weatherService services.WeatherService
}

func NewWeatherUseCase(zipCodeRepo repository.ZipCodeRepository, weatherService services.WeatherService) *WeatherUseCase {
	return &WeatherUseCase{
		zipCodeRepo:    zipCodeRepo,
		weatherService: weatherService,
	}
}

func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (map[string]float64, error) {
	location, err := u.zipCodeRepo.GetLocationByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	var query string
	if location.Bairro != "" && location.Localidade != "" {
		query = fmt.Sprintf("%s, %s", location.Bairro, location.Localidade)
	} else if location.Localidade != "" {
		query = location.Localidade
	} else {
		return nil, errors.New("localização não encontrada")
	}

	return u.weatherService.GetWeather(query)
}
