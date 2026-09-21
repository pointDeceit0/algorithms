package leetcode


func resultArray(nums []int, k int) []int64 {
	// Time ~ O(nk)
	res := make([]int64, k)
	dp := make([]int64, k)

	for _, v := range nums {
		res[v % k]++
		tmpDp := make([]int64, k)
		for r := 0; r < k; r++ {
			tmpDp[(r * v) % k] += dp[r]
		}

		tmpDp = dp

		for r := 0; r < k; r++ {
			res[r] += dp[r]
		}
	}

	return res
}
