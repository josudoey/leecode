package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("maxProfit",
	func(nums []int, expected int) {
		Expect(maxProfit(nums)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{7, 1, 5, 3, 6, 4},
		5,
	),
	Entry("Case 2",
		[]int{7, 6, 4, 3, 1},
		0,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
