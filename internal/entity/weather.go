package entity

import (
	"fmt"
	"strconv"

	"github.com/pedrogutierresbr/lab-weather-api-pos-goexpert/pkg/temperature"
)

type Weather struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

type WeatherDetails struct {
	Temperature float64
}

type WeatherResponse struct {
	WeatherResponse WeatherDetails
}

func NewWeather(celsius, fahrenheit, kelvin float64) *Weather {
	return &Weather{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	}
}

func (w *Weather) ConvertTemperatures(celsius float64) {
	w.Celsius, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", celsius), 64)
	w.Fahrenheit, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temperature.ConvertCelsiusToFahrenheit(celsius)), 64)
	w.Kelvin, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temperature.ConvertCelsiusToKelvin(celsius)), 64)
}
