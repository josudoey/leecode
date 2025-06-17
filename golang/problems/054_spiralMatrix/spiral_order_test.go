package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("spiralOrder",
	func(matrix [][]int, expected []int) {
		Expect(spiralOrder(matrix)).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	),
	Entry("Case 2",
		[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
