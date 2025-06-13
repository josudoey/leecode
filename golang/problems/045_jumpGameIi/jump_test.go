package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("jump",
	func(nums []int, expected int) {
		Expect(jump(nums)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{2, 3, 1, 1, 4},
		2,
	),
	Entry("Case 2",
		[]int{2, 3, 0, 1, 4},
		2,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
