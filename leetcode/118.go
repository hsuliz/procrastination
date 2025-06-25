package leetcode

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		currentRow := i + 1
		row := make([]int, currentRow)
		for j := 0; j < currentRow; j++ {
			row[j] = 1
		}
		triangle[i] = row
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}
