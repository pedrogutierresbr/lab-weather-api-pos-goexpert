package web_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/infra/web"
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

func (m *MockCEPRepository) GetCEP(cep string) ([]byte, error) {
	args := m.Called(cep)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCEPRepository) ConvertResponse(cepData []byte) (*entity.CEP, error) {
	args := m.Called(cepData)
	return args.Get(0).(*entity.CEP), args.Error(1)
}

type MockWeatherRepository struct {
	mock.Mock
}

func (m *MockWeatherRepository) GetWeather(city, apiKey string) ([]byte, error) {
	args := m.Called(city, apiKey)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockWeatherRepository) ConvertResponse(weatherData []byte) (*entity.WeatherResponse, error) {
	args := m.Called(weatherData)
	return args.Get(0).(*entity.WeatherResponse), args.Error(1)
}

func (m *MockWeatherRepository) ConvertToAllWeathers(weatherData *entity.WeatherResponse) (*entity.Weather, error) {
	args := m.Called(weatherData)
	return args.Get(0).(*entity.Weather), args.Error(1)
}

func TestWebCepHandler_Get(t *testing.T) {
	mockCEPRepo := new(MockCEPRepository)
	mockWeatherRepo := new(MockWeatherRepository)
	handler := web.NewHTTPHandler(mockCEPRepo, mockWeatherRepo, "test-api-key")

	r := chi.NewRouter()
	r.Get("/cep/{cep}", handler.Get)

	t.Run("success", func(t *testing.T) {
		cep := "12220-780"
		cepData := []byte(`{"cep": "12220-780", "localidade": "São José dos Campos", "uf": "SP"}`)
		cepResponse := &entity.CEP{CEP: cep, Localidade: "São José dos Campos", UF: "SP"}
		weatherData := []byte(`{"current": {"temp_c": 25.3}}`)
		weatherResponse := &entity.WeatherResponse{}
		weather := &entity.Weather{Celsius: 25.3}

		mockCEPRepo.On("IsValidCEP", cep).Return(true)
		mockCEPRepo.On("GetCEP", cep).Return(cepData, nil)
		mockCEPRepo.On("ConvertResponse", cepData).Return(cepResponse, nil)
		mockWeatherRepo.On("GetWeather", cepResponse.Localidade, "mock-api-key").Return(weatherData, nil)
		mockWeatherRepo.On("ConvertResponse", weatherData).Return(weatherResponse, nil)
		mockWeatherRepo.On("ConvertToAllWeathers", weatherResponse).Return(weather, nil)

		req := httptest.NewRequest(http.MethodGet, "/cep/"+cep, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response entity.Weather
		json.NewDecoder(w.Body).Decode(&response)
		assert.Equal(t, 25.3, response.Celsius)
	})
}
