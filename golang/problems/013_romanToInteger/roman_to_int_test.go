package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("romanToInt",
	func(s string, expected int) {
		Expect(romanToInt(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"III",
		3,
	),
	Entry("Case 2",
		"LVIII",
		58,
	),
	Entry("Case 3",
		"MCMXCIV",
		1994,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
