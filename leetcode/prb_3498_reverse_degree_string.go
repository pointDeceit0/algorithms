package leetcode


func ReverseDegree(s string) int {

	var res int
	for i := 0; i < len(s); i++ {
		res += int(123 - s[i]) * (i + 1)
	}

	return res
}