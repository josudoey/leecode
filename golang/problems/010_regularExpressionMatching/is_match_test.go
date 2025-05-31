package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("isMatch",
	func(s string, p string, expected bool) {
		Expect(isMatch(s, p)).To(Equal(expected))
	},
	Entry("Case 1",
		"11", "a",
		false,
	),
	Entry("Case 2",
		"aa", "a*",
		true,
	),
	Entry("Case 3",
		"ab", ".*",
		true,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
