package _542

import (
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	var queue [][]int
	rowSize := len(mat)
	colSize := len(mat[0])

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if mat[i][j] == 1 {
				mat[i][j] = math.MaxInt32
			} else {
				queue = append(queue, []int{i, j})
			}
		}
	}
	rowVector := []int{-1, 0, 1, 0}
	colVector := []int{0, 1, 0, -1}
	for len(queue) != 0 {
		position := queue[0]
		row := position[0]
		col := position[1]
		for i := 0; i < 4; i++ {
			nextRow := row + rowVector[i]
			nextCol := col + colVector[i]
			if nextRow < 0 || nextRow > rowSize-1 ||
				nextCol < 0 || nextCol > colSize-1 {
				continue
			}
			if mat[nextRow][nextCol] > mat[row][col]+1 {
				mat[nextRow][nextCol] = mat[row][col] + 1
				queue = append(queue, []int{nextRow, nextCol})
			}
		}
		queue = append(queue[:0], queue[1:]...)
	}

	return mat
}
