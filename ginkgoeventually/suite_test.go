package ginkgofailhandler

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEventually(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eventually Suite")
}

var _ = Describe("D1", func() {
	It("It1", func() {
		i := 0
		Eventually(func(g Gomega) {
			i += 1
			time.Sleep(5 * time.Second)
			g.Expect(true).To(BeFalse())
		}).Should(Succeed())

		Expect(i).To(Equal(1))
	})
})
