package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("twoSum",
	func(nums []int, target int, expected []int) {
		Expect(twoSum(nums, target)).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 2},
	),
	Entry("Case 2",
		[]int{2, 3, 4},
		6,
		[]int{1, 3},
	),
	Entry("Case 3",
		[]int{-1, 0},
		-1,
		[]int{1, 2},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
