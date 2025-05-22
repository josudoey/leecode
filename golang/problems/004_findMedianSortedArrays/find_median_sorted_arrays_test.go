package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("findMedianSortedArrays",
	func(nums1 []int, nums2 []int, expected float64) {
		Expect(findMedianSortedArrays(nums1, nums2)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{1, 3},
		[]int{2},
		2.0,
	),
	Entry("Case 2",
		[]int{1, 2},
		[]int{3, 4},
		2.5,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
