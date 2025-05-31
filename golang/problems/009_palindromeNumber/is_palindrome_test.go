package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("isPalindrome",
	func(x int, expected bool) {
		Expect(isPalindrome(x)).To(Equal(expected))
	},
	Entry("Case 1",
		121,
		true,
	),
	Entry("Case 2",
		-121,
		false,
	),
	Entry("Case 3",
		10,
		false,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
