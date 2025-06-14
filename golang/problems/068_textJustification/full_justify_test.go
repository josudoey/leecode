package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("fullJustify",
	func(words []string, maxWidth int, expected []string) {
		Expect(fullJustify(words, maxWidth)).To(Equal(expected))
	},
	Entry("Case 1",
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
	),
	Entry("Case 2",
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	),
	Entry("Case 3",
		[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20,
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
