package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("reverse",
	func(x int, expected int) {
		Expect(reverse(x)).To(Equal(expected))
	},
	Entry("Case 1",
		123,
		321,
	),
	Entry("Case 2",
		-123,
		-321,
	),
	Entry("Case 3",
		120,
		21,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
