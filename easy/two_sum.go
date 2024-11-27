package main

import "fmt"

func main(){
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
}

func twoSum(nums []int, target int) []int {
    mapped := make(map[int]int)

    for i, v := range nums {
		needed := target - v
        ind, matched := mapped[needed]

		if matched && ind != i {
			return []int{ind, i}
		}
        mapped[v] = i
        
    }

    return []int{}
}