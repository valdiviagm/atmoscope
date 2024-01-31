// weather_report_validator.go

package weather_report

//import "errors"

type WeatherReportValidator struct {}

func NewWeatherReportValidator() *WeatherReportValidator{
	return &WeatherReportValidator{}
}

func (wrv *WeatherReportValidator) Validate( wr WeatherReport ) error {
	

	return nil
}

