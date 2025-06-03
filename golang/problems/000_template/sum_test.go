package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("sum",
	func(a int, b int, expected int) {
		Expect(sum(a, b)).To(Equal(expected))
	},
	Entry("Case 1",
		1,
		2,
		3,
	),
	Entry("Case 2",
		4,
		5,
		9,
	),
	Entry("Case 3",
		6,
		7,
		13,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
