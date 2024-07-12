package nested

import (
	. "github.com/onsi/ginkgo/v2"
)


var _ = BeforeSuite(func() {
	By("Nested BeforeSuite")
})

var _ = AfterSuite(func() {
	By("Nested AfterSuite")
})

var _ = Describe("D1", func() {
	It("It1", func() {
		By("It1")
	})
})
