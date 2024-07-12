package gomegaformat

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGomegaFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gomega Format Suite")
}

var _ = Describe("D1", func() {
	It("It1", func() {
		Expect(false).To(BeTrue(), "Formatted message with string[%s]", "foo")
	})
})
