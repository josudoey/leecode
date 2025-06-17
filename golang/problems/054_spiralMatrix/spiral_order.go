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
	resultLength := m * n
	if resultLength == 1 {
		return []int{matrix[0][0]}
	}

	result := make([]int, 0, resultLength)
	topIndex, bottomLimit := 0, len(matrix)
	leftIndex, rightLimit := 0, len(matrix[0])
	for leftIndex < rightLimit {
		for i := leftIndex; i < rightLimit; i++ {
			result = append(result, matrix[topIndex][i])
		}
		topIndex++

		for i := topIndex; i < bottomLimit; i++ {
			result = append(result, matrix[i][rightLimit-1])
		}
		rightLimit--

		if leftIndex == rightLimit || topIndex == bottomLimit {
			break
		}

		for i := rightLimit - 1; i >= leftIndex; i-- {
			result = append(result, matrix[bottomLimit-1][i])
		}
		bottomLimit--

		for i := bottomLimit - 1; i >= topIndex; i-- {
			result = append(result, matrix[i][leftIndex])
		}
		leftIndex++

		if leftIndex == rightLimit || topIndex == bottomLimit {
			break
		}
	}

	return result
}
