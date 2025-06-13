package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("reverseWords",
	func(s string, expected string) {
		Expect(reverseWords(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"the sky is blue",
		"blue is sky the",
	),
	Entry("Case 2",
		"  hello world  ",
		"world hello",
	),
	Entry("Case 3",
		"a good   example",
		"example good a",
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
