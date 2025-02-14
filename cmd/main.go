package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/configs"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/repository"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/services"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/usecase"
)

func main() {
	cfg := configs.GetConfig()

	zipCodeRepo := repository.NewZipCodeRepository()
	weatherService := services.NewWeatherService(cfg.WeatherAPIKey)
	weatherUseCase := usecase.NewWeatherUseCase(zipCodeRepo, weatherService)

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		zipCode := r.URL.Query().Get("cep")
		if zipCode == "" {
			http.Error(w, "O parâmetro 'cep' é obrigatório", http.StatusBadRequest)
			return
		}

		weather, err := weatherUseCase.GetWeatherByZipCode(zipCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weather)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
