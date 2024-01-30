// main.go
package main

import (
    "fmt"
    "github.com/valdiviagm/atmoscope/src/weather_report"
)



func main() {

    // Initialize WeatherReport instance
    wr := weather_report.NewWeatherReport(1.0, 2.0, 3.0)

    // Use getter methods to access values
    fmt.Println("Temperature:", wr.GetTemperature())
    fmt.Println("Humidity:", wr.GetHumidity())
    fmt.Println("WindSpeed:", wr.GetWindSpeed())

    // Use setter methods to update values
    wr.SetTemperature(4.0)
    wr.SetHumidity(5.0)
    wr.SetWindSpeed(6.0)
    
    // Use getter methods to access updated values
    fmt.Println("Temperature:", wr.GetTemperature())
    fmt.Println("Humidity:", wr.GetHumidity())
    fmt.Println("WindSpeed:", wr.GetWindSpeed())
  
}

