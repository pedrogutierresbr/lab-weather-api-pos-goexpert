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

type WeatherResponse struct {
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		TZID      string  `json:"tz_id"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PrecipMm   float64 `json:"precip_mm"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
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
