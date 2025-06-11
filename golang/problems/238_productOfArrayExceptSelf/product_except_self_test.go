package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("productExceptSelf",
	func(nums []int, expected []int) {
		Expect(productExceptSelf(nums)).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[]int{1, 2, 3, 4},
		[]int{24, 12, 8, 6},
	),
	Entry("Case 2",
		[]int{-1, 1, 0, -3, 3},
		[]int{0, 0, 9, 0, 0},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
