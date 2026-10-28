package leetcode


func rec3882(grid *[][]int, filled *[][][]int, i, j, val int) int {

	cur_val := val ^ (*grid)[i][j]

	if (*filled)[i][j][cur_val] == 1 {
		return cur_val
	}

	(*filled)[i][j][cur_val] = 1

	if len(*grid) - 1 == i && len((*grid)[0]) - 1 == j {
		return cur_val
	}

	if i + 1 < len(*grid) {
		rec3882(grid, filled, i + 1, j, cur_val)
	}
	if j + 1 < len((*grid)[0]) {
		rec3882(grid, filled, i, j + 1, cur_val)
	}

	return -1
}


func minCost(grid [][]int) int {

	var N, M, MAX_VAL int = len(grid), len(grid[0]), 1024

	filled := make([][][]int, N)
	for i := range filled {
		filled[i] = make([][]int, M)
		for j := range filled[i] {
			filled[i][j] = make([]int, MAX_VAL)
		}
	}

	rec3882(&grid, &filled, 0, 0, 0)

	for i, v := range filled[N - 1][M - 1] {
		if v == 1 {
			return i
		}
	}
	return 0
}