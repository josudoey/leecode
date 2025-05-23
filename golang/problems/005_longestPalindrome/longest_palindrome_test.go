package code

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("longestPalindrome",
	func(s string, expected string) {
		Expect(longestPalindrome(s)).To(Equal(expected))
	},
	Entry("Case 1",
		"babad",
		"bab",
	),
	Entry("Case 2",
		"cbbd",
		"bb",
	),
	Entry("Case 3",
		"a",
		"a",
	),
	Entry("Case 4",
		"bb",
		"bb",
	),
	Entry("Case 5",
		"abb",
		"bb",
	),
	Entry("Case 6",
		"aacabdkacaa",
		"aca",
	),
	Entry("case 7",
		"jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel",
		"sknks",
	),
)

func TestCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Code")
}
