package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("isPalindrome",
	func(s string, expected bool) {
		Expect(isPalindrome(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"A man, a plan, a canal: Panama",
		true,
	),
	Entry("Case 2",
		"race a car",
		false,
	),
	Entry("Case 3",
		" ",
		true,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
