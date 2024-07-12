package ginkgonestedsuite

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/zeeke/go-samples/ginkgonestedsuite/nested"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Before/After Suite")
}


var _ = BeforeSuite(func() {
	By("Parent BeforeSuite")
})

var _ = AfterSuite(func() {
	By("Parent AfterSuite")
})