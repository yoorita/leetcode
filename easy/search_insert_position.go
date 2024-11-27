package main

func searchInsert(nums []int, target int) int {
    for i, v := range nums {
		if v == target || target < v{
			return i
		}
	}
    return len(nums)
}