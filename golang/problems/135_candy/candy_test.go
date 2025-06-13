package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("candy",
	func(ratings []int, expected int) {
		Expect(candy(ratings)).To(Equal(expected))
	},
	Entry("Case 1",
		[]int{1, 0, 2},
		5,
	),
	Entry("Case 2",
		[]int{1, 2, 2},
		4,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
