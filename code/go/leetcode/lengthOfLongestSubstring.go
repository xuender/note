package leetcode

import "strings"

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}
func lengthOfLongestSubstring2(s string) int {
	l := len(s)
	ret := 0
	for i := 0; i < l; i++ {
		m := make(map[byte]int)
		for f := i; f < l; f++ {
			m[s[f]] = 1
			if len(m) <= f-i {
				if f-i > ret {
					ret = f - i
				}
				m = nil
				break
			}
		}
		if m != nil && len(m) > ret {
			ret = len(m)
		}
	}
	return ret
}
