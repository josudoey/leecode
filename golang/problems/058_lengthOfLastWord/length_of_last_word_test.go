package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("lengthOfLastWord",
	func(s string, expected int) {
		Expect(lengthOfLastWord(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"Hello World",
		5,
	),
	Entry("Case 2",
		"   fly me   to   the moon  ",
		4,
	),
	Entry("Case 3",
		"luffy is still joyboy",
		6,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
