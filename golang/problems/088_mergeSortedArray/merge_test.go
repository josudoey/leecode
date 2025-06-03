package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("merge",
	func(nums1 []int, m int, nums2 []int, n int, expected []int) {
		merge(nums1, m, nums2, n)
		Expect(nums1).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[]int{1, 2, 3, 0, 0, 0},
		3,
		[]int{2, 5, 6},
		3,
		[]int{1, 2, 2, 3, 5, 6},
	),
	Entry("Case 2",
		[]int{1},
		1,
		[]int{},
		0,
		[]int{1},
	),
	Entry("Case 2",
		[]int{0},
		0,
		[]int{1},
		1,
		[]int{1},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
