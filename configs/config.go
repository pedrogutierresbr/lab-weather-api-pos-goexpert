package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherAPIKey string
}

var config *Config

func LoadConfig() {
	// Tenta carregar o arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env (pode estar ausente)")
	}

	config = &Config{
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
	}

	if config.WeatherAPIKey == "" {
		log.Fatal("Chave da API do WeatherAPI não configurada (WEATHER_API_KEY)")
	}

	log.Println("Configurações carregadas com sucesso")
}

func GetConfig() *Config {
	if config == nil {
		log.Fatal("Configurações não carregadas. Certifique-se de chamar LoadConfig antes de GetConfig.")
	}
	return config
}
