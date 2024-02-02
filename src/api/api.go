package api

import (
	"io"
	"net/http"
	"encoding/json"
	"github.com/valdiviagm/atmoscope/src/weather_report"
)

// API represents the API client with data and error channels
type API struct {
	DataChannel chan interface{}
	ErrorChannel chan error
}

func NewAPI() *API {
	return &API{
		DataChannel: make( chan interface{} ),
		ErrorChannel: make( chan error ),
	} 
}

func (api *API) FetchWeatherReport( url string ) {

	response, err := http.Get(url)
	if err != nil {
		api.ErrorChannel <- err
		return
	}
	defer response.Body.Close()
        
	
	body, err := io.ReadAll(response.Body) 
	if err != nil {
		api.ErrorChannel <- err
		return
	}

	var data weather_report.LocationInfo

	if err := json.Unmarshal(body, &data); err != nil {
		api.ErrorChannel <- err
		return
	} else {

		wr := weather_report.NewWeatherReport(
			data.Current.Temperature2m,
			float64(data.Current.RelativeHumidity2m),
			data.Current.WindSpeed10m)
	
		api.DataChannel <- wr

	}
	




}
