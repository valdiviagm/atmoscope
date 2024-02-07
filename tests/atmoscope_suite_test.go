package atmoscope_test

import (
	"testing"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
	"net/http"
	"net/http/httptest"
	"github.com/valdiviagm/atmoscope/src/weather_report"
   	"github.com/valdiviagm/atmoscope/src/api"

)

func TestAtmoscope(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Atmoscope Suite")
}

var _ = Describe("WeatherReportValidator", func() {

	var wrv *weather_report.WeatherReportValidator

	BeforeEach( func() {
		wrv = weather_report.NewWeatherReportValidator()
	})

	It("should validate a valid WeatherReport", func() {

		wr := weather_report.WeatherReport{}
	
		err:= wrv.Validate( wr )
		Expect(err).To(BeNil())
	
	})

	It("should validate a non-valid 'humidity'", func(){
		
		wr := weather_report.NewWeatherReport(0,120,0)
	
		err:= wrv.Validate( wr )
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Humidity is out of valid range"))
	})

	It("should validate a non-valid 'windSpeed'", func(){
		
		wr := weather_report.NewWeatherReport(0,0,-10)
	
		err:= wrv.Validate( wr )
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Wind speed must be non-negative"))
	})


	It("should validate a non-valid 'temperature'", func(){
		
		wr := weather_report.NewWeatherReport(-300,0,0)
	
		err:= wrv.Validate( wr )
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Temperture must obey the laws of physics"))
	
	})

	// More test cases to be added later
})

var _ = Describe("WeatherReport", func() {

	var wr weather_report.WeatherReport

	BeforeEach( func() {
		wr = weather_report.NewWeatherReport(0,0,0)
	})

	Context("SetTemperature", func() {

		It("should correctly set 'temperature'", func() {

			wr.SetTemperature(25)
			Expect(wr.GetTemperature()).To(Equal(25.0))
		
		})

	})

	Context("SetHumidity", func() {

		It("should correctly set 'humidity'", func() {

			wr.SetHumidity(50)
			Expect(wr.GetHumidity()).To(Equal(50.0))
		
		})

	})

	Context("SetWindSpeed", func() {

		It("should correctly set 'windSpeed'", func() {

			wr.SetWindSpeed(75)
			Expect(wr.GetWindSpeed()).To(Equal(75.0))
		
		})

	})




	// More test cases to be added later
})


func StartMockAPIServer() (*httptest.Server) {

	mockResponse := `{
		"latitude": 52.52,
		"longitude": 13.419998,
		"generationtime_ms": 0.07200241088867188,
		"utc_offset_seconds": 0,
		"timezone": "GMT",
		"timezone_abbreviation": "GMT",
		"elevation": 38.0,
		"current_units": {
			"time": "iso8601",
			"interval": "seconds",
			"temperature_2m": "°C",
			"wind_speed_10m": "km/h",
			"relative_humidity_2m": "%"
		},
		"current": {
			"time": "2024-02-06T23:15",
			"interval": 900,
			"temperature_2m": 9.2,
			"wind_speed_10m": 26.4,
			"relative_humidity_2m": 81
		}
	}`

    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
	switch r.URL.Path {
        case "/v1/forecast":
            // Handle the regular request (e.g., for "/v1/forecast")
            
	    w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            fmt.Fprintln(w, mockResponse)
        case "/test-timeout":
	    	fmt.Println("whitin /test-timeout...")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusGatewayTimeout)
		fmt.Fprintln(w, `{"error": "Timeout"}`)
        default:
            // Handle other requests (e.g., not found)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintln(w, `{"error": "Not found"}`)
        }
    }))
}

var _ = Describe("API Integration", func() {
    // Variables for the mock API server
    var server *httptest.Server

    // Create a new instance of your custom API
    API := api.NewAPI()

    // Start the mock API server before each test
    BeforeEach(func() {
        // Start the mock API server and get its URL
        server = StartMockAPIServer()
    })

    // Close the mock API server after each test
    AfterEach(func() {
        server.Close()
    })

    It("should return WeatherReport data", func() {
        // Construct the URL to the test server

        // Start fetching data from the API using a goroutine

        apiURL := fmt.Sprintf("%s/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,wind_speed_10m", server.URL)

	go API.FetchWeatherReport(apiURL)

        select {
        case Data := <-API.DataChannel:
            // Unmarshal the JSON response data into a WeatherReport struct
           
	    wr := Data.(weather_report.WeatherReport) 

            Expect(wr.GetTemperature()).To(Equal(9.2))
            Expect(wr.GetHumidity()).To(Equal(81.0))
            Expect(wr.GetWindSpeed()).To(Equal(26.4))

        case Err := <-API.ErrorChannel:
            Expect(Err).To(BeNil())

        }
    })


    It("should return error on bad API request", func() {
        // Construct the URL to the test server
        apiURL := fmt.Sprintf("%s/", server.URL)
	
        // Start fetching data from the API using a goroutine
        go API.FetchWeatherReport(apiURL)

        select {

        case Err := <-API.ErrorChannel:
            Expect(Err).To(Equal(errors.New("LocationInfo struct was not modified after unmarshaling")))
	}

    })




})


