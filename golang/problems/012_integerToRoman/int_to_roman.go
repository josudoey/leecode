package code

import "strings"

// ref https://leetcode.com/problems/integer-to-roman/

// Example 1:
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

// Example 2:
// Input: height = [1,1]
// Output: 1

var romanSymbols = [][]string{
	{
		"",     // 0
		"I",    // 1
		"II",   // 2
		"III",  // 3
		"IV",   // 4
		"V",    // 5
		"VI",   // 6
		"VII",  // 7
		"VIII", // 8
		"IX",   // 9
	}, {
		"",     // 0
		"X",    // 10
		"XX",   // 20
		"XXX",  // 30
		"XL",   // 40
		"L",    // 50
		"LX",   // 60
		"LXX",  // 70
		"LXXX", // 80
		"XC",   // 90
	}, {
		"",     // 0
		"C",    // 100
		"CC",   // 200
		"CCC",  // 300
		"CD",   // 400
		"D",    // 500
		"DC",   // 600
		"DCC",  // 700
		"DCCC", // 800
		"CM",   // 900
	}, {
		"",    // 0
		"M",   // 1000
		"MM",  // 2000
		"MMM", // 3000
	},
}

func intToRoman(num int) string {
	var symbols []string
	for i := 0; num != 0; i++ {
		digit := num % 10
		symbols = append(symbols, romanSymbols[i][digit])
		num /= 10
	}

	var b strings.Builder
	for i := len(symbols) - 1; i >= 0; i-- {
		b.WriteString(symbols[i])
	}
	return b.String()
}
