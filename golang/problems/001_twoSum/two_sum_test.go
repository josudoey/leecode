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
		[]int{0, 1},
	),
	Entry("Case 2",
		[]int{3, 2, 4},
		6,
		[]int{1, 2},
	),
	Entry("Case 3",
		[]int{3, 3},
		6,
		[]int{0, 1},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
