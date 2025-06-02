package code

import (
	"sort"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func normalizeResult(nums [][]int) [][]int {
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

var _ = DescribeTable("fourSum",
	func(nums []int, target int, expected [][]int) {
		Expect(normalizeResult(fourSum(nums, target))).
			To(BeComparableTo(normalizeResult(expected)))
	},
	Entry("Case 1",
		[]int{1, 0, -1, 0, -2, 2},
		0,
		[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
	),
	Entry("Case 2",
		[]int{2, 2, 2, 2, 2},
		8,
		[][]int{{2, 2, 2, 2}},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
