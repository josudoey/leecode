package code

// ref https://leetcode.com/problems/rotate-image/

// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]

// Example 2:
// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

func rotate(matrix [][]int) {
	n := len(matrix)
	iLimit := n / 2

	if n%2 == 1 {
		iLimit++
	}

	for i := 0; i < iLimit; i++ {
		invI := n - i - 1
		for j := 0; j < n/2; j++ {
			invJ := n - j - 1
			matrix[i][j], matrix[invJ][i],
				matrix[invI][invJ], matrix[j][invI] =
				matrix[invJ][i], matrix[invI][invJ],
				matrix[j][invI], matrix[i][j]

		}
	}
}
