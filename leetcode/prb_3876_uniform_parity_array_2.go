package leetcode

func uniformArray(nums1 []int) bool {

	var min_odd, min_even int

	for _, v := range nums1 {
		if v % 2 == 1 && (min_odd == 0 || v < min_odd) {
			min_odd = v
		} else {
			min_even = v
		}
	}

	if min_odd > min_even && min_even != 0 {
		return false
	}

	return true
}