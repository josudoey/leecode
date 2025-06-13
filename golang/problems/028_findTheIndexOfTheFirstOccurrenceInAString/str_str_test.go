package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("strStr",
	func(haystack string, needle string, expected int) {
		Expect(strStr(haystack, needle)).To(Equal(expected))
	},
	Entry("Case 1",
		"sadbutsad",
		"sad",
		0,
	),
	Entry("Case 2",
		"leetcode",
		"leeto",
		-1,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
