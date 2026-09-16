package leetcode


func MaxPalindromes(s string, k int) int {

	var res int
	left_start := 0

	for right := 1; right < len(s); right++ {
		for left := left_start; left < right; left++ {
			for left < right && s[left] == s[right] {
				left++; right--
			}
			if s[left] != s[right] {
				break
			} else {
				res++
			}
		}

	}

	return res
}