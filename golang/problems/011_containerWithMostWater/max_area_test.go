package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("maxArea",
	func(height []int, expected int) {
		Expect(maxArea(height)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	),
	Entry("Case 2",
		[]int{1, 1},
		1,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
