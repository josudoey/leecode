package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("rotate",
	func(nums []int, k int, expected []int) {
		rotate(nums, k)
		Expect(nums).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	),
	Entry("Case 2",
		[]int{-1, -100, 3, 99},
		2,
		[]int{3, 99, -1, -100},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
