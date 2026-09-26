package leetcode

import "strings"

func evaluate(s string, knowledge [][]string) string {
	// This isn't the most efficient solution, there's no necessity:
	// 1. to create tmp builder
	// 2. to write byte at each step, it could be done by slices
	// Time  ~ O(n + m)
	// Space ~ O(n + m) -- tmp could be grown to n size

	d := make(map[string]string)
	res := &strings.Builder{}

	for _, v := range knowledge {
		d[v[0]] = v[1]
	}

	for i := 0; i < len(s); i++ {
		if s[i] != '(' && s[i] != ')' {
			res.WriteByte(s[i])
		} else {
			i++
			tmp := &strings.Builder{}
			for ; s[i] != ')'; i++ {
				tmp.WriteByte(s[i])
			}
			if val, ok := d[tmp.String()]; ok {
				res.WriteString(val)
			} else {
				res.WriteByte('?')
			}
		}
	}

	return res.String()
}