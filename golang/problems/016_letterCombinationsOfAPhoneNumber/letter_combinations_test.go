package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("letterCombinations",
	func(digits string, expected []string) {
		Expect(letterCombinations(digits)).To(Equal(expected))
	},
	Entry("Case 1",
		"23",
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	),
	Entry("Case 2",
		"",
		[]string{},
	),
	Entry("Case 3",
		"2",
		[]string{"a", "b", "c"},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
