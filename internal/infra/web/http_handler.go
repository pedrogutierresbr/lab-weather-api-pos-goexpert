package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
)

type WebCepHandler struct {
	CEPRepository     entity.CEPRepositoryInterface
	WeatherRepository entity.WeatherRepositoryInterface
	ApiKey            string
}

func NewHTTPHandler(cepRepository entity.CEPRepositoryInterface, weatherRepository entity.WeatherRepositoryInterface, apiKey string) *WebCepHandler {
	return &WebCepHandler{
		CEPRepository:     cepRepository,
		WeatherRepository: weatherRepository,
		ApiKey:            apiKey,
	}
}

func (h *WebCepHandler) Get(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	if !h.CEPRepository.IsValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cepData, err := h.CEPRepository.GetCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	cepResponse, err := h.CEPRepository.ConvertResponse(cepData)
	if err != nil {
		http.Error(w, "can not convert zipcode", http.StatusInternalServerError)
		return
	}

	weatherData, err := h.WeatherRepository.GetWeather(cepResponse.Localidade, h.ApiKey)
	if err != nil {
		http.Error(w, "can not get weather", http.StatusInternalServerError)
		return
	}

	weatherResponse, err := h.WeatherRepository.ConvertResponse(weatherData)
	if err != nil {
		http.Error(w, "can not convert weather data", http.StatusInternalServerError)
		return
	}

	weather, err := h.WeatherRepository.ConvertToAllWeathers(weatherResponse)
	if err != nil {
		http.Error(w, "can not convert weather response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
