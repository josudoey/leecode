package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = ""
// Output: []

// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]

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
