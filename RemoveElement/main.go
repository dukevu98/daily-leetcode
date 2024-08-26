package main

import "fmt"

func main() {
	//fmt.Println(removeElement([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	fmt.Println(nums)
	return left
}
