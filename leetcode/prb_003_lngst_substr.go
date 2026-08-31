package leetcode

func lengthOfLongestSubstring(s string) int {

    d := make(map[byte]int)
	cur_str := ""
	max_len := 0

	for i := 0; i < len(s); i++ {
		if _, ok := d[s[i]]; ok {
			j := 0
			for ; j < len(cur_str) && d[s[i]] > 0; j++ {
				d[cur_str[j]] -= 1
			}
			cur_str = cur_str[j:]
			d[s[i]] += 1
		} else {
			d[s[i]] = 1
		}
		cur_str += string(s[i])
		max_len = max(max_len, len(cur_str))
	}

	return max_len
}