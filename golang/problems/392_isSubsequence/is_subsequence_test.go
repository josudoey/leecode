package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("isSubsequence",
	func(s string, t string, expected bool) {
		Expect(isSubsequence(s, t)).To(Equal(expected))
	},
	Entry("Case 1",
		"abc",
		"ahbgdc",
		true,
	),
	Entry("Case 2",
		"axc",
		"ahbgdc",
		false,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
