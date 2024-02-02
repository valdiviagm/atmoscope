// weather_report.go
package weather_report

type CurrentUnits struct {
	Time               string `json:"time"`
	Interval           string `json:"interval"`
	Temperature2m      string `json:"temperature_2m"`
	WindSpeed10m       string `json:"wind_speed_10m"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
}

type CurrentData struct {
	Time            string  `json:"time"`
	Interval        int     `json:"interval"`
	Temperature2m   float64 `json:"temperature_2m"`
	WindSpeed10m    float64 `json:"wind_speed_10m"`
	RelativeHumidity2m int	`json:"relative_humidity_2m"`
}

type LocationInfo struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	GenerationTimeMs    float64 `json:"generationtime_ms"`
	UtcOffsetSeconds    int     `json:"utc_offset_seconds"`
	Timezone            string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation           float64 `json:"elevation"`
	CurrentUnits        CurrentUnits `json:"current_units"`
	Current             CurrentData `json:"current"`
}

type WeatherReport struct {
	temperature 	float64
	humidity 	float64
	windSpeed 	float64
}


// NewWeatherReport creates a new WeatherReport instance with the specified values.
func NewWeatherReport(temperature, humidity, windSpeed float64) WeatherReport {
    return WeatherReport{
        temperature: temperature,
        humidity:    humidity,
        windSpeed:   windSpeed,
    }
}

// GetTemperature returns the temperature value.
func (wr *WeatherReport) GetTemperature() float64 {
    return wr.temperature
}

// GetHumidity returns the humidity value.
func (wr *WeatherReport) GetHumidity() float64 {
    return wr.humidity
}

// GetWindSpeed returns the windSpeed value.
func (wr *WeatherReport) GetWindSpeed() float64 {
    return wr.windSpeed
}

// SetTemperature sets the temperature value.
func (wr *WeatherReport) SetTemperature(newTemperature float64) {
    wr.temperature = newTemperature
}

// SetHumidity sets the humidity value.
func (wr *WeatherReport) SetHumidity(newHumidity float64) {
    wr.humidity = newHumidity
}

// SetWindSpeed sets the windSpeed value.
func (wr *WeatherReport) SetWindSpeed(newWindSpeed float64) {
    wr.windSpeed = newWindSpeed
}

