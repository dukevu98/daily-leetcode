package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{11, 15, 2, 7}, 9))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := numMap[target-nums[i]]; ok {
			return []int{j, i}
		}
		numMap[nums[i]] = i
	}
	return result
}
