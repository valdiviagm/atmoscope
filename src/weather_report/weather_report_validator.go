// weather_report_validator.go

package weather_report

import "errors"

type WeatherReportValidator struct {}

func NewWeatherReportValidator() *WeatherReportValidator{
	return &WeatherReportValidator{}
}

func (wrv *WeatherReportValidator) Validate( wr WeatherReport ) error {
	
	if wr.GetHumidity() < 0.0 || wr.GetHumidity() > 100.0 {
		return errors.New("Humidity is out of valid range")
	} 

	if wr.GetWindSpeed() < 0.0 {
		return errors.New("Wind speed must be non-negative")
	} 

	if wr.GetTemperature() < -215.15 {
		return errors.New("Temperture must obey the laws of physics")
	} 

	return nil
}

