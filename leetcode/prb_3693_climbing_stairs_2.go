package leetcode


import ("math")


func rec3693(minCost, costs *[]int, i int) int {

	if i == len(*costs) {return 0}

	if (*minCost)[i] != -1 {return (*minCost)[i]}

	m := math.MaxInt64
	for j := 1; j < 4; j++ {
		if i + j <= len(*costs) {
			m = min(m, rec3693(minCost, costs, i + j) + (*costs)[i + j - 1] + j * j)
		}
	}

	(*minCost)[i] = m

	return (*minCost)[i]
}


func climbStairs(n int, costs []int) int {
 
	minCost := make([]int, n)

	for i := range minCost {
		minCost[i] = -1
	}

	return rec3693(&minCost, &costs, 0)
}