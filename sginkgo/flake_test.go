package sginkgo

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flake Suite")
}

var _ = Describe("D1", FlakeAttempts(1), func() {
	It("It1", FlakeAttempts(1), func() {
		Expect(false).To(BeTrue())
	})
})
