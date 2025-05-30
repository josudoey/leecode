package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("myAtoi",
	func(s string, expected int) {
		Expect(myAtoi(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"42",
		42,
	),
	Entry("Case 2",
		" -042",
		-42,
	),
	Entry("Case 3",
		"1337c0d3",
		1337,
	),
	Entry("Case 4",
		"0-1",
		0,
	),
	Entry("Case 5",
		"words and 987",
		0,
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
