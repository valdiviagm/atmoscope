// main.go
package main

import (
    "fmt"
    "time"
    "github.com/valdiviagm/atmoscope/src/weather_report"
    "github.com/valdiviagm/atmoscope/src/api"
)

func main() {

    API := api.NewAPI()

    go API.FetchWeatherReport("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,wind_speed_10m")
    
    timeout := time.After( 5 * time.Second )
    
    select{
    	case Data := <- API.DataChannel:

	    wr := Data.(weather_report.WeatherReport) 
	    
	    fmt.Println("Temperature:", wr.GetTemperature() )
	    fmt.Println("Humidity:", wr.GetHumidity() )
	    fmt.Println("Wind Speed:", wr.GetWindSpeed() )

   	case Err := <- API.ErrorChannel:
		fmt.Println(Err)
	case <- timeout:
		fmt.Println("No response received after timeout reached")
    }
    
    return

}


