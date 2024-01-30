package atmoscope_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAtmoscope(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Atmoscope Suite")
}
