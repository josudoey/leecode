package code

import "strings"

// ref https://leetcode.com/problems/zigzag-conversion/

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

func convert(s string, numRows int) string {
	length := len(s)
	numCells := (length / numRows) + 1
	tableIndexes := make([][]int, numRows)
	for rowIndex := range tableIndexes {
		tableIndexes[rowIndex] = make([]int, 0, numCells)
	}

	oneZigzagLength := numRows
	if numRows > 2 {
		oneZigzagLength += numRows - 2
	}

	for i := range length {
		zigzagRowIndex := (i % oneZigzagLength)
		reverseIndex := oneZigzagLength - zigzagRowIndex
		if zigzagRowIndex > reverseIndex {
			zigzagRowIndex = reverseIndex
		}

		tableIndexes[zigzagRowIndex] = append(tableIndexes[zigzagRowIndex], i)
	}

	result := strings.Builder{}
	for rowIndex := range tableIndexes {
		for _, index := range tableIndexes[rowIndex] {
			result.WriteByte(s[index])
		}
	}

	return result.String()
}
