package code

// ref https://leetcode.com/problems/valid-sudoku/

// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true

// Example 2:
// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

func isValidSudoku(board [][]byte) bool {
	rowCounts := make([]int, 81)
	columnCounts := make([]int, 81)
	gridCounts := make([]int, 81)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			charCode := compressChar(board[i][j])
			if charCode < 0 {
				continue
			}

			rowIndex := 9*i + charCode
			if rowCounts[rowIndex] == 1 {
				return false
			}
			rowCounts[rowIndex]++

			columnIndex := 9*j + charCode
			if columnCounts[columnIndex] == 1 {
				return false
			}
			columnCounts[columnIndex]++

			gridIndex := 9*((i/3)*3+j/3) + charCode
			if gridCounts[gridIndex] == 1 {
				return false
			}
			gridCounts[gridIndex]++
		}
	}

	return true
}

func compressChar(char byte) int {
	if '1' <= char && char <= '9' {
		return int(char - '1')
	}

	return -1
}
