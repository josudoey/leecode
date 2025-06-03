package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("removeDuplicates",
	func(nums []int, expected []int) {
		Expect(removeDuplicates(nums)).To(Equal(len(expected)))

		distinctNums := nums[:len(expected)]
		Expect(distinctNums).To(BeComparableTo(expected))

	},
	Entry("Case 1",
		[]int{1, 1, 2},
		[]int{1, 2},
	),
	Entry("Case 2",
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		[]int{0, 1, 2, 3, 4},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
