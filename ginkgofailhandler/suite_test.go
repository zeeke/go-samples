package ginkgofailhandler

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(func(message string, callerSkip ...int) {
		By("Fail Handler")

		By("CurrentSpecReport().LeafNodeText: " + CurrentSpecReport().LeafNodeText)
		By("CurrentSpecReport().Failed(): " + fmt.Sprintf("%t", CurrentSpecReport().Failed()))

		Fail(message, callerSkip...)
	})
	RunSpecs(t, "Fail Hanlder Suite")
}

var _ = Describe("D1", func() {
	It("It1", func() {
		Expect(false).To(BeTrue())
	})
})
