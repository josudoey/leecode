package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("convert",
	func(s string, numRows int, expected string) {
		Expect(convert(s, numRows)).To(Equal(expected))
	},
	Entry("Case 1",
		"PAYPALISHIRING",
		3,
		"PAHNAPLSIIGYIR",
	),
	Entry("Case 2",
		"PAYPALISHIRING",
		4,
		"PINALSIGYAHRPI",
	),
	Entry("Case 3",
		"A",
		1,
		"A",
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
