package code

// ref https://leetcode.com/problems/spiral-matrix/

// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]

// Example 2:
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	if m*n == 1 {
		return []int{matrix[0][0]}
	}
	result := make([]int, m*n)
	gone := make([]interface{}, m*n)
	moveMatrix := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	currentState := 0

	i, j := 0, 0
	for index := range result {
		result[index] = matrix[i][j]
		gone[i*n+j] = struct{}{}
		offset := moveMatrix[currentState]
		nextI, nextJ := i+offset[0], j+offset[1]

		for nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n {
			currentState = (currentState + 1) % 4
			offset := moveMatrix[currentState]
			nextI, nextJ = i+offset[0], j+offset[1]
		}

		if gone[(nextI*n)+nextJ] == struct{}{} {
			currentState = (currentState + 1) % 4
			offset := moveMatrix[currentState]
			nextI, nextJ = i+offset[0], j+offset[1]
		}
		i, j = nextI, nextJ
	}

	return result

}
