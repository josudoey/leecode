package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("canCompleteCircuit",
	func(gas []int, cost []int, expected int) {

		Expect(canCompleteCircuit(gas, cost)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5, 1, 2},
		3,
	),
	Entry("Case 2",
		[]int{2, 3, 4},
		[]int{3, 4, 3},
		-1,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
