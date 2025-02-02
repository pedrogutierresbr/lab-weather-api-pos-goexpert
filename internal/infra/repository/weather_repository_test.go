package repository_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/infra/repository"
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

func TestGetWeather(t *testing.T) {
	weatherJson := getWeatherJSON()
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(weatherJson)),
				Header:     make(http.Header),
			}, nil
		},
	}

	repo := repository.NewWeatherRepository(mockClient)
	response, err := repo.GetWeather("Uberaba", "mock-api-key")

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestConvertResponse(t *testing.T) {
	weatherJson := getWeatherJSON()
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(weatherJson)),
				Header:     make(http.Header),
			}, nil
		},
	}

	repo := repository.NewWeatherRepository(mockClient)
	weatherResponse, err := repo.ConvertResponse(weatherJson)

	assert.Nil(t, err)
	assert.NotNil(t, weatherResponse)
	assert.Equal(t, "Uberaba", weatherResponse.Location.Name)
	assert.Equal(t, 25.3, weatherResponse.Current.TempC)
}

func TestConvertToAllWeathers(t *testing.T) {
	weatherJson := getWeatherJSON()

	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(weatherJson)),
				Header:     make(http.Header),
			}, nil
		},
	}

	repo := repository.NewWeatherRepository(mockClient)
	weatherResponse, _ := repo.ConvertResponse(weatherJson)

	weather, err := repo.ConvertToAllWeathers(weatherResponse)

	assert.NoError(t, err)
	assert.IsType(t, weather, &entity.Weather{})

	assert.Equal(t, fmt.Sprintf("%.2f", weather.Celsius*1.8+32), fmt.Sprintf("%.2f", weather.Fahrenheit))
	assert.Equal(t, fmt.Sprintf("%.2f", weather.Celsius+273), fmt.Sprintf("%.2f", weather.Kelvin))
}
