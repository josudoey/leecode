package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("minSubArrayLen",
	func(target int, nums []int, expected int) {
		Expect(minSubArrayLen(target, nums)).To(Equal(expected))
	},
	Entry("Case 1",
		7,
		[]int{2, 3, 1, 2, 4, 3},
		2,
	),
	Entry("Case 2",
		4,
		[]int{1, 4, 4},
		1,
	),
	Entry("Case 3",
		11,
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		0,
	),
	Entry("Case 4",
		213,
		[]int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
		8,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
