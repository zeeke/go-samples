package multipleskips

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEventually(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eventually Suite")
}

var _ = Describe("D1", func() {
	It("A", func() {
		Fail("It1 failure")
	})

	It("B", func() {
		Fail("It1 failure")
	})

	It("C", func() {
		Fail("It1 failure")
	})
})
