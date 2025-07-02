package leetcode

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	ogColor := image[sr][sc]
	if ogColor == color {
		return image
	}

	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		if r >= len(image) || r < 0 || c >= len(image[0]) || c < 0 {
			return
		}
		if image[r][c] != ogColor {
			return
		}
		image[r][c] = color

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	dfs(sr, sc)
	return image
}
