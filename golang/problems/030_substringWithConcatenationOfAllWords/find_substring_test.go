package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("findSubstring",
	func(s string, words []string, expected []int) {
		Expect(findSubstring(s, words)).To(BeComparableTo(expected))
	},
	Entry("Case 1",
		"barfoothefoobarman",
		[]string{"foo", "bar"},
		[]int{0, 9},
	),
	Entry("Case 2",
		"wordgoodgoodgoodbestword",
		[]string{"word", "good", "best", "word"},
		[]int{},
	),
	Entry("Case 3",
		"barfoofoobarthefoobarman",
		[]string{"bar", "foo", "the"},
		[]int{6, 9, 12},
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
