package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("hIndex",
	func(nums []int, expected int) {
		Expect(hIndex(nums)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{3, 0, 6, 1, 5},
		3,
	),
	Entry("Case 2",
		[]int{1, 3, 1},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
