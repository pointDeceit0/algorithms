package leetcode


import ("slices")


type Point struct {
	i int
	k int
}


func ABS(v int) int {
	if v >= 0 {return v}
	return -v
}


func insertSorted(arr []int, val int) []int {
	idx, _ := slices.BinarySearch(arr, val)
	return slices.Insert(arr, idx, val)
}


func recMinOperations(ops *[]int, passed *map[Point]int, i, k, x, kGoal int) int {
	if v, ok := (*passed)[Point{i, k}]; ok && v != -1 {
        return v + (*ops)[i]
    } else if v == -1 {
        return -1
    }

	if k == kGoal {return (*ops)[i]}

	if len(*ops) < i + x {return -1}


	m := -1
	for j := i + x; j < len(*ops); j++ {
		ret := recMinOperations(ops, passed, j, k + 1, x, kGoal)
		if ret == -1 {
			continue
		} else if m == -1 || ret < m {
			m = ret
		}
	}
	(*passed)[Point{i, k}] = m

    if m == -1 {return -1}
	
	return (*ops)[i] + m
}


func minOperations(nums []int, x int, k int) int64 {
	n := len(nums)
	ops := []int{}

	// ~O(n xlogx)
	// TLE
	for i := x; i <= n; i++ {

		tmp := []int{}
		for j := i - x; j < i; j++ {
			tmp = insertSorted(tmp, nums[j])
		}

		tmpOps := 0
		for _, v := range tmp {
			tmpOps += ABS(tmp[x / 2] - v)
		}

		ops = append(ops, tmpOps)
	}

	passed := make(map[Point]int)

	m := int(1e16)
	for j := 0; j <= n - x; j++ {
		ret := recMinOperations(&ops, &passed, j, 1, x, k)
		if ret == -1 {
			continue
		} else if m == -1 || ret < m {
			m = ret
		}
	}

	return int64(m)
}