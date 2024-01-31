package atmoscope_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	// Import the weather_report package
	"github.com/valdiviagm/atmoscope/src/weather_report"

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







