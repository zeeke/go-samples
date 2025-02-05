package ginkgopointerequals

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fail Hanlder Suite")
}

var _ = Describe("D1", func() {
	It("It1", func() {
		var a *string
		var b *string

		s1 := "xxx"
		s2 := "xxx"
		a = &s1
		b = &s2
		Expect(a).To(Equal(b))
	})
})
