package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/internal/entity"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherRepository struct {
	client HTTPClient
}

func NewWeatherRepository(client HTTPClient) *WeatherRepository {
	return &WeatherRepository{
		client: client,
	}
}

func (w *WeatherRepository) GetWeather(localidade string, apiKey string) ([]byte, error) {
	localidade = strings.Replace(localidade, " ", "%20", -1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, localidade), nil)
	if err != nil {
		log.Printf("Falha ao criar requisição HTTP: %v", err)
		return nil, err
	}

	res, err := w.client.Do(req)
	if err != nil {
		log.Printf("Falha ao realizar requisição HTTP: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	ctxErr := ctx.Err()
	if ctxErr != nil {
		log.Printf("Timeout ao realizar requisição HTTP: %v", ctxErr)
		return nil, err
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Falha ao ler resposta HTTP: %v", err)
		return nil, err
	}

	if strings.Contains(string(response), "Invalid API key") {
		return nil, err
	}

	return response, nil
}

func (w *WeatherRepository) ConvertResponse(weatherData []byte) (*entity.WeatherResponse, error) {
	var weatherResponse entity.WeatherResponse
	err := json.Unmarshal(weatherData, &weatherResponse)
	if err != nil {
		log.Printf("Falha ao converter resposta HTTP: %v", err)
		return nil, err
	}

	return &weatherResponse, nil
}

func (w *WeatherRepository) ConvertToAllWeathers(weatherResponse *entity.WeatherResponse) (*entity.Weather, error) {
	weather := entity.Weather{}
	weather.ConvertTemperatures(weatherResponse.Current.TempC)

	return &weather, nil
}
