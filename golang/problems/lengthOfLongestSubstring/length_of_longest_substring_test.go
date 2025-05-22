package lengthOfLongestSubstring

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("lengthOfLongestSubstring",
	func(s string, expected int) {
		Expect(lengthOfLongestSubstring(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"abcabcbb",
		3,
	),
	Entry("Case 2",
		"bbbbb",
		1,
	),
	Entry("Case 3",
		"pwwkew",
		3,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
