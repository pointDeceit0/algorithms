package leetcode


type state struct { i, j, s int }


func recHasValidPath(s int, visited *map[state]bool, grid *[][]byte, i, j int) bool {
	// Brute force ~ O(2^(n + m))
	// With memory optimization ~ O(mn(m + n))

	initS := s

	if (*visited)[state{i, j, s}] {
		return false
	}

	if (*grid)[i][j] == 40 {
		s++
	} else if (*grid)[i][j] == 41 && s > 0 {
		s--
	} else {
		(*visited)[state{i, j, s}] = true
		return false
	}

	res := false || s == 0 && len(*grid) - 1 == i && len((*grid)[0]) - 1 == j

	if i + 1 < len(*grid) {
		res = res || recHasValidPath(s, visited, grid, i + 1, j)
	}
	if !res && j + 1 < len((*grid)[0]) {
		res = res || recHasValidPath(s, visited, grid, i, j + 1)
	}

    if !res {
        (*visited)[state{i, j, initS}] = true
    }
	
	return res
}


func hasValidPath(grid [][]byte) bool {
    if (len(grid) + len(grid[0]) - 1) % 2 != 0 {return false}

	visited := make(map[state]bool)

	return recHasValidPath(0, &visited, &grid, 0, 0)
}