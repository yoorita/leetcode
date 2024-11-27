package main

func removeElement(nums []int, val int) int {
    var keeped int

	for _,v := range nums {
		if v != val {
			nums[keeped] = v
			keeped++
		}
	}

	return keeped
}