package leetcode

func twoSum1(nums []int, target int) []int {
	for i, n := range nums {
		for f, m := range nums {
			if i != f && n+m == target {
				return []int{i, f}
			}
		}
	}
	return []int{-1, -1}
}
func twoSum2(nums []int, target int) []int {
	old := make([]int, len(nums))
	for i, n := range nums {
		for t := 0; t < i; t++ {
			if old[t] == n {
				return []int{t, i}
			}
		}
		old[i] = target - n
	}
	return []int{-1, -1}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			return []int{value, i}
		}
		m[k] = i
	}
	return []int{-1, -1}
}
