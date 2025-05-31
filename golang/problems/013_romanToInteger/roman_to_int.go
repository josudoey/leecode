package code

// ref https://leetcode.com/problems/roman-to-integer/

// Example 1:
// Input: s = "III"
// Output: 3
// Explanation: III = 3.

// Example 2:
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.

// Example 3:
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

var romanSymbolMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
	"CC":   200,
	"CCC":  300,
	"CD":   400,
	"D":    500,
	"DC":   600,
	"DCC":  700,
	"DCCC": 800,
	"CM":   900,
	"M":    1000,
	"MM":   2000,
	"MMM":  3000,
}

func romanToInt(s string) int {
	var (
		result = 0
		head   = ""
		value  = 0
	)

	for _, c := range s {
		symbol := head + string(c)
		v, ok := romanSymbolMap[symbol]
		if ok {
			head = symbol
			value = v
			continue
		}

		result += value
		head = string(c)
		value = romanSymbolMap[head]
	}

	result += value
	return result
}
