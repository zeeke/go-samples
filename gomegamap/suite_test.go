package gomegamap

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type label string
type Metric map[label]label

func TestGomega(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gomega Map Suite")
}

var _ = Describe("D1", func() {
	It("It1", func() {
		var x Metric = Metric{}
		x["node"] = "xxx"
		x["foo"] = "bar"

		Expect(x).To(And(
			HaveKey(label("node")),
			HaveKey("foo"),
		))
	})
})
