package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("intToRoman",
	func(num int, expected string) {
		Expect(intToRoman(num)).To(Equal(expected))
	},
	Entry("Case 1",
		3749,
		"MMMDCCXLIX",
	),
	Entry("Case 2",
		58,
		"LVIII",
	),
	Entry("Case 3",
		1994,
		"MCMXCIV",
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
