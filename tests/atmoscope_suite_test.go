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

	// More test cases to be added later
})








