package leetcode

import (
	"strings"
)


func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


func rec(j int, s string) (int, string) {

	tmp := &strings.Builder{}
	i := j
	for ; i < len(s); i++ {
		switch s[i] {
		case '(':
			tmp.WriteString(s[j:i])
			beg, r := rec(i + 1, s)
			i, j = beg, beg + 1
			tmp.WriteString(r)
		case ')':
			tmp.WriteString(s[j:i])
			return i, Reverse(tmp.String())
		}
	}
	tmp.WriteString(s[i:])
	return 0, tmp.String()
}


func ReverseParentheses(s string) string {

	_, res := rec(0, s)
	return res
}
