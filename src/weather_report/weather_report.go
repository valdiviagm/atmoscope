// weather_report.go
package weather_report

type WeatherReport struct {
    temperature float64
    humidity    float64
    windSpeed   float64
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

