package entity

type CEPRepositoryInterface interface {
	IsValidCEP(string) bool
	GetCEP(string) ([]byte, error)
	ConvertResponse([]byte) (*CEP, error)
}

type WeatherRepositoryInterface interface {
	GetWeather(string, string) ([]byte, error)
	ConvertResponse([]byte) (*WeatherResponse, error)
	ConvertToAllWeathers(*WeatherResponse) (*Weather, error)
}
