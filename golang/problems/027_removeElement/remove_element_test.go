package code

import (
	"sort"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("removeElement",
	func(nums []int, val int, expected []int) {
		Expect(removeElement(nums, val)).To(Equal(len(expected)))

		removedSlice := nums[:len(expected)]
		sort.Ints(removedSlice)
		sort.Ints(expected)
		Expect(removedSlice).To(BeComparableTo(expected))

	},
	Entry("Case 1",
		[]int{3, 2, 2, 3},
		3,
		[]int{2, 2},
	),
	Entry("Case 2",
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
		[]int{0, 1, 4, 0, 3},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
