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
	// Verifica se a API key está definida no ambiente
	configs.LoadConfig()
	cfg := configs.GetConfig()

	// Inicializa os repositórios e serviços
	zipCodeRepo := repository.NewZipCodeRepository()
	weatherService := services.NewWeatherService(cfg.WeatherAPIKey)
	weatherUseCase := usecase.NewWeatherUseCase(zipCodeRepo, weatherService)

	// Define a rota para a API
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		zipCode := r.URL.Query().Get("cep") // Altere para capturar "cep" corretamente
		if zipCode == "" {
			http.Error(w, "O parâmetro 'cep' é obrigatório", http.StatusBadRequest)
			return
		}

		// Obtém o clima pelo CEP
		weather, err := weatherUseCase.GetWeatherByZipCode(zipCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retorna a resposta como JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weather)
	})

	// Inicializa o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão para desenvolvimento local
	}
	log.Printf("Servidor rodando na porta %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
