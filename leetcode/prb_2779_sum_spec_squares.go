package leetcode


func sumOfSquares(nums []int) int {

	n := len(nums)
	res := 0

	for i, v := range nums {
		if n % (i + 1) == 0 {
			res += v * v
		}
	}
    
	return res
}