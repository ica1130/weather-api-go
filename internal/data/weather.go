package data

type Weather struct {
	Address           string            `json:"address"`
	Description       string            `json:"description"`
	CurrentConditions CurrentConditions `json:"currentConditions"`
}

type CurrentConditions struct {
	Temp       float64 `json:"temp"`
	FeelsLike  float64 `json:"feelslike"`
	Humidity   float64 `json:"humidity"`
	Dew        float64 `json:"dew"`
	Snow       float64 `json:"snow"`
	SnowDepth  float64 `json:"snowdepth"`
	WindGust   float64 `json:"windgust"`
	WindSpeed  float64 `json:"windspeed"`
	Pressure   float64 `json:"pressure"`
	Visibility float64 `json:"visibility"`
	Sunrise    string  `json:"sunrise"`
	Sunset     string  `json:"sunset"`
}
