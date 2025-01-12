package temperature

func ConvertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func ConvertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
