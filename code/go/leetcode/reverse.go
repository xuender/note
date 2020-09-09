package leetcode

import "strconv"

func reverse(x int) int {
	str := strconv.Itoa(x)
	tmp := []byte(str)
	if x < 0 {
		tmp = tmp[1:]
	}
	var strBytes []byte
	length := len(tmp)
	for i := 0; i < length; i++ {
		strBytes = append(strBytes, tmp[length-i-1])
	}
	str = string(strBytes)
	max := 1<<31 - 1
	min := ^max
	if x < 0 {
		r, _ := strconv.Atoi("-" + str)
		if r > max || r < min {
			return 0
		}
		return r
	}
	r, _ := strconv.Atoi(str)
	if r > max || r < min {
		return 0
	}
	return r
}
