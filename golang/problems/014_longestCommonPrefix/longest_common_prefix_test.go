package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("longestCommonPrefix",
	func(strs []string, expected string) {
		Expect(longestCommonPrefix(strs)).To(Equal(expected))
	},
	Entry("Case 1",
		[]string{"flower", "flow", "flight"},
		"fl",
	),
	Entry("Case 2",
		[]string{"dog", "racecar", "car"},
		"",
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
