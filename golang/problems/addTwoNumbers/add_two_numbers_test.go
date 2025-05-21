package addTwoNumbers

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("addTwoNumbers",
	func(input1 []int, input2 []int, expected []int) {
		l1 := newListNode(input1)
		l2 := newListNode(input2)
		result := addTwoNumbers(l1, l2)
		Expect(newIntSlice(result)).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[]int{2, 4, 3},
		[]int{5, 6, 4},
		[]int{7, 0, 8},
	),
	Entry("Case 2",
		[]int{0},
		[]int{0},
		[]int{0},
	),
	Entry("Case 3",
		[]int{9, 9, 9, 9, 9, 9, 9},
		[]int{9, 9, 9, 9},
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
