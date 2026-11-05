package leetcode


func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	for _, v1 := range arr1 {
		flag := true
		for _, v2 := range arr2 {
			if ABS(v1 - v2) <= d {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	return cnt
}