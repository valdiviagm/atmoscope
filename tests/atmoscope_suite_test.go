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

	// More test cases to be added later
})








