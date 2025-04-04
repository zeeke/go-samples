package ginkgofailhandler

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fail In Message Descriptor Suite")
}

var _ = Describe("D1", func() {
	It("It1", func() {
		Expect(false).To(BeTrue(), func() string {
			Expect(false).To(BeTrue())
			return "fff"
		})
	})
})
