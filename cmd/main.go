package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/configs"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/infra/repository"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/infra/web"
	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/infra/web/webserver"
)

func ConfigureServer() *webserver.WebServer {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	webserver := webserver.NewWebServer(configs.WebServerPort)

	cepRepo := repository.NewCEPRepository()
	weatherRepo := repository.NewWeatherRepository(&http.Client{})

	openWeatherMapAPIKey := configs.OpenWeatherAPIKey
	if openWeatherMapAPIKey == "" {
		log.Fatal("Please, provide the OPEN_WEATHERMAP_API_KEY environment variable. Make sure you provide a valid API key, otherwise it will not be possible to get and convert weather data.")
	}

	webCEPHandler := web.NewHTTPHandler(cepRepo, weatherRepo, openWeatherMapAPIKey)
	webserver.AddHandler("GET /cep/{cep}", webCEPHandler.Get)

	return webserver
}

func main() {
	webserver := ConfigureServer()
	fmt.Println("Starting web server on port", ":8080")
	webserver.Start()
}
