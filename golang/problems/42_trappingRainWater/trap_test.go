package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("trap",
	func(height []int, expected int) {
		Expect(trap(height)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		6,
	),
	Entry("Case 2",
		[]int{4, 2, 0, 3, 2, 5},
		9,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
