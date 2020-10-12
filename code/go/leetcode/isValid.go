package leetcode

func isValid(s string) bool {
	cs := []rune{}
	for _, i := range s {
		switch i {
		case '(', '[', '{':
			cs = append(cs, i)
		case ')':
			if len(cs) < 1 || cs[len(cs)-1] != '(' {
				return false
			}
			cs = cs[:len(cs)-1]
		case ']':
			if len(cs) < 1 || cs[len(cs)-1] != '[' {
				return false
			}
			cs = cs[:len(cs)-1]
		case '}':
			if len(cs) < 1 || cs[len(cs)-1] != '{' {
				return false
			}
			cs = cs[:len(cs)-1]
		}
	}
	return len(cs) == 0
}
