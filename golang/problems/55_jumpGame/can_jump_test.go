package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("canJump",
	func(nums []int, expected bool) {
		Expect(canJump(nums)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{2, 3, 1, 1, 4},
		true,
	),
	Entry("Case 2",
		[]int{3, 2, 1, 0, 4},
		false,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
