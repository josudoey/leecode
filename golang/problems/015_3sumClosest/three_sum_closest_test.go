package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("threeSumClosest",
	func(nums []int, target int, expected int) {
		Expect(threeSumClosest(nums, target)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{-1, 2, 1, -4}, 1,
		2,
	),
	Entry("Case 2",
		[]int{0, 0, 0}, 1,
		0,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
