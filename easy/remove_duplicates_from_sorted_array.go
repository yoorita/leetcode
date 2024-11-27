package main

func removeDuplicates(nums []int) int {
    uniqueId := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueId] {
			nums[uniqueId+1] = nums[i]
			uniqueId++
		}
	}
	return uniqueId+1
}