package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("minWindow",
	func(s string, t string, expected string) {
		Expect(minWindow(s, t)).To(Equal(expected))
	},
	Entry("Case 1",
		"ADOBECODEBANC",
		"ABC",
		"BANC",
	),
	Entry("Case 2",
		"a",
		"a",
		"a",
	),
	Entry("Case 3",
		"a",
		"aa",
		"",
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
