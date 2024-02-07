package api

import (
	"io"
	"errors"
	"reflect"
	"net/http"
	"context"
	"time"
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

func (api *API) FetchWeatherReport(url string) {
    // Create a context with a timeout of 3 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()


    // Make the HTTP request within the context
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        api.ErrorChannel <- err
        return
    }

    response, err := http.DefaultClient.Do(req)
    if err != nil {
        // Check if the error is due to a timeout
        if ctx.Err() == context.DeadlineExceeded {
            api.ErrorChannel <- ctx.Err()
        } else {
	    api.ErrorChannel <- err
	}
        return
    }
    defer response.Body.Close()

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        api.ErrorChannel <- err
        return
    }

    // Unmarshal the response body into the appropriate struct
    var data weather_report.LocationInfo
   
    // Create a copy of the original struct
    var dataBeforeUnmarshal = data
    
    // Unmarshal the response body into the appropriate struct
    if err := json.Unmarshal(body, &data); err != nil {
        api.ErrorChannel <- err
        return
    }

    // Check if the struct was modified
    if reflect.DeepEqual(dataBeforeUnmarshal, data) {
        api.ErrorChannel <- errors.New("LocationInfo struct was not modified after unmarshaling")
	return
    }

    // Process the data and send it to the DataChannel
    wr := weather_report.NewWeatherReport(
        data.Current.Temperature2m,
        float64(data.Current.RelativeHumidity2m),
        data.Current.WindSpeed10m,
    )
    api.DataChannel <- wr

}
