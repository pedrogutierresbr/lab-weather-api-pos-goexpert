package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	celsius := 25.0
	fahrenheit := 77.0
	kelvin := 298.15

	weather := NewWeather(celsius, fahrenheit, kelvin)

	assert.NotNil(t, weather, "A instância de Weather não deveria ser nula")
	assert.Equal(t, celsius, weather.Celsius, "Os valores de Celsius não são iguais")
	assert.Equal(t, fahrenheit, weather.Fahrenheit, "Os valores de Fahrenheit não são iguais")
	assert.Equal(t, kelvin, weather.Kelvin, "Os valores de Kelvin não são iguais")
}

func TestMakeTemperatureConversions(t *testing.T) {
	weather := NewWeather(1, 1, 1)
	weather.ConvertTemperatures(27.1)
	assert.Equal(t, weather.Celsius, 27.1)
	assert.Equal(t, weather.Fahrenheit, weather.Celsius*1.8+32)
	assert.Equal(t, weather.Kelvin, weather.Celsius+273)
}
