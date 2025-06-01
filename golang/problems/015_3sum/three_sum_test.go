package code

import (
	"sort"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func normalizeThreeSumResult(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] < nums[j][0] {
			return false
		}
		if nums[i][1] < nums[j][1] {
			return false
		}
		return nums[i][2] < nums[j][2]
	})

	return nums
}

var _ = DescribeTable("threeSum",
	func(nums []int, expected [][]int) {
		Expect(normalizeThreeSumResult(threeSum(nums))).
			To(BeComparableTo(normalizeThreeSumResult(expected)))
	},
	Entry("Case 1",
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	),
	Entry("Case 2",
		[]int{0, 1, 1},
		[][]int{},
	),
	Entry("Case 3",
		[]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10},
		[][]int{{-10, 5, 5}, {-5, 0, 5}, {-4, 2, 2}, {-3, -2, 5}, {-3, 1, 2}, {-2, 0, 2}},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
