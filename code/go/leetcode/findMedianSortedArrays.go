package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	length := l1 + l2
	a := length / 2
	nums3 := make([]int, a+1)

	i, j, m := 0, 0, 0
	for i < l1 && j < l2 {
		if m > a {
			break
		}
		if nums1[i] > nums2[j] {
			nums3[m] = nums2[j]
			j++
		} else {
			nums3[m] = nums1[i]
			i++
		}
		m++
	}

	for ; m < a+1 && j < l2; j++ {
		nums3[m] = nums2[j]
		m++
	}
	for ; m < a+1 && i < l1; i++ {
		nums3[m] = nums1[i]
		m++
	}

	if length%2 == 0 {
		return float64(nums3[a-1]+nums3[a]) / 2.0
	}
	return float64(nums3[a])
}
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	all := append(nums1, nums2...)
	sort.Ints(all)
	l := len(all)
	if l%2 == 1 {
		return float64(all[l/2])
	}
	return float64(all[l/2-1]+all[l/2]) / 2
}
