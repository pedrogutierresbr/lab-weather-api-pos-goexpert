package usecase

import (
	"errors"
	"strings"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
)

type WeatherInputDTO struct {
	Localidade string
	ApiKey     string
}

type WeatherOutputDTO struct {
	Celsius   float64 `json:"celsius"`
	Farenheit float64 `json:"farenheit"`
	Kelvin    float64 `json:"kelvin"`
}

type GetWeatherUsecase struct {
	WeatherRepository entity.WeatherRepositoryInterface
}

func NewWeatherUseCase(weatherRepository entity.WeatherRepositoryInterface) *GetWeatherUsecase {
	return &GetWeatherUsecase{
		WeatherRepository: weatherRepository,
	}
}

func (w *GetWeatherUsecase) Execute(input WeatherInputDTO) (WeatherOutputDTO, error) {
	if input.Localidade == "" {
		return WeatherOutputDTO{}, errors.New("localidade não informada")
	}

	if input.ApiKey == "" {
		return WeatherOutputDTO{}, errors.New("apiKey não informada")
	}

	weatherResponse, err := w.WeatherRepository.GetWeather(input.Localidade, input.ApiKey)
	if err != nil || strings.Contains(string(weatherResponse), "city not found") {
		return WeatherOutputDTO{}, errors.New("falha ao recuperar a temperatura")
	}

	weather, err := w.WeatherRepository.ConvertResponse(weatherResponse)
	if err != nil {
		return WeatherOutputDTO{}, errors.New("falha ao converter a temperatura")
	}

	weatherConverted, err := w.WeatherRepository.ConvertToAllWeathers(weather)
	if err != nil {
		return WeatherOutputDTO{}, errors.New("falha ao converter para demais unidades de temperatura")
	}

	dto := WeatherOutputDTO{
		Celsius:   weatherConverted.Celsius,
		Farenheit: weatherConverted.Fahrenheit,
		Kelvin:    weatherConverted.Kelvin,
	}

	return dto, nil
}
