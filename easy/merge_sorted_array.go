package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i, v := range nums2 {
		nums1[m+i] = v
	}
	sort.IntSlice(nums1).Sort()
}