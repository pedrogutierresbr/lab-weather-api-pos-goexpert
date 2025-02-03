package usecase_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/infra/repository"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/usecase"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func getWeatherJSON() []byte {
	return []byte(
		`{
			"location": {
				"name": "Uberaba",
				"region": "Minas Gerais",
				"country": "Brazil",
				"lat": -19.75,
				"lon": -47.9167,
				"tz_id": "America/Sao_Paulo",
				"localtime_epoch": 1738526095,
				"localtime": "2025-02-02 16:54"
			},
			"current": {
				"last_updated_epoch": 1738525500,
				"last_updated": "2025-02-02 16:45",
				"temp_c": 25.3,
				"temp_f": 77.5,
				"is_day": 1,
				"condition": {
				"text": "Overcast",
				"icon": "//cdn.weatherapi.com/weather/64x64/day/122.png",
				"code": 1009
				},
				"wind_mph": 10.5,
				"wind_kph": 16.9,
				"wind_degree": 25,
				"wind_dir": "NNE",
				"pressure_mb": 1011.0,
				"pressure_in": 29.85,
				"precip_mm": 0.07,
				"precip_in": 0.0,
				"humidity": 79,
				"cloud": 50,
				"feelslike_c": 27.8,
				"feelslike_f": 82.0,
				"windchill_c": 22.8,
				"windchill_f": 73.1,
				"heatindex_c": 25.0,
				"heatindex_f": 77.0,
				"dewpoint_c": 20.0,
				"dewpoint_f": 68.0,
				"vis_km": 10.0,
				"vis_miles": 6.0,
				"uv": 0.8,
				"gust_mph": 14.4,
				"gust_kph": 23.1
			}
		}`,
	)
}

func NewMockClient(json_type string) *MockClient {
	var weatherJson = getWeatherJSON()

	if json_type == "invalid" {
		return &MockClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewReader([]byte(``))),
					Header:     make(http.Header),
				}, errors.New("fail to get weather")
			},
		}
	}

	return &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			mockResponse := weatherJson
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(mockResponse))),
				Header:     make(http.Header),
			}, nil
		},
	}
}

func TestGetWeather(t *testing.T) {
	mockClientWithValidJSON := NewMockClient("valid")
	weatherRepositoryWithValidJSON := repository.NewWeatherRepository(mockClientWithValidJSON)
	getWeatherWithValidJSON := usecase.NewWeatherUseCase(weatherRepositoryWithValidJSON)

	mockClientWithInvalidJSON := NewMockClient("invalid")
	weatherRepositoryWithInvalidJSON := repository.NewWeatherRepository(mockClientWithInvalidJSON)
	getWeatherWithInvalidJSON := usecase.NewWeatherUseCase(weatherRepositoryWithInvalidJSON)

	t.Run("temperatura válida", func(t *testing.T) {
		getWeatherDTO := usecase.WeatherInputDTO{
			Localidade: "Uberaba",
			ApiKey:     "mock-api-key",
		}

		weatherOutput, err := getWeatherWithValidJSON.Execute(getWeatherDTO)
		assert.NoError(t, err)
		assert.IsType(t, usecase.WeatherOutputDTO{}, weatherOutput)
	})

	t.Run("localidade não informada", func(t *testing.T) {
		getWeatherDTO := usecase.WeatherInputDTO{
			ApiKey: "mock-api-key",
		}

		weatherOutput, err := getWeatherWithValidJSON.Execute(getWeatherDTO)
		assert.EqualError(t, err, "localidade não informada")
		assert.IsType(t, usecase.WeatherOutputDTO{}, weatherOutput)
	})

	t.Run("apiKey não informada", func(t *testing.T) {
		getWeatherDTO := usecase.WeatherInputDTO{
			Localidade: "Uberaba",
		}

		weatherOutput, err := getWeatherWithValidJSON.Execute(getWeatherDTO)
		assert.EqualError(t, err, "apiKey não informada")
		assert.IsType(t, usecase.WeatherOutputDTO{}, weatherOutput)
	})

	t.Run("falha ao recuperar a temperatura", func(t *testing.T) {
		getWeatherDTO := usecase.WeatherInputDTO{
			Localidade: "São Pauloo",
			ApiKey:     "mock-api-key",
		}

		weatherOutput, err := getWeatherWithInvalidJSON.Execute(getWeatherDTO)
		assert.EqualError(t, err, "falha ao recuperar a temperatura")
		assert.Equal(t, weatherOutput.Celsius, float64(0))
		assert.Equal(t, weatherOutput.Fahrenheit, float64(0))
		assert.Equal(t, weatherOutput.Kelvin, float64(0))
	})
}
