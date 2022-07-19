package main

import (
	"sort"
)

func main() {
	var arr, arr2, arr3, arr4 []int
	arr = append(arr, 1, 2, 3, 4, 0, 0)
	arr2 = append(arr2, 1, 7)
	arr3 = append(arr, 1, 2, 3, 4, 0, 0)
	arr4 = append(arr2, 1, 7)
	merge(arr, 4, arr2, 2)
	mergeII(arr3, 4, arr4, 2)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m])
	nums2 = append(nums2[:n])
	for i := 0; i < len(nums2); i++ {
		nums1 = append(nums1, nums2[i])
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}

func mergeII(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}
