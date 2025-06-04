package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("majorityElement",
	func(nums []int, expected int) {
		Expect(majorityElement(nums)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{3, 2, 3},
		3,
	),
	Entry("Case 2",
		[]int{2, 2, 1, 1, 1, 2, 2},
		2,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
