package main

import "fmt"

func main() {
	var a = sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(a)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, len(nums)-1
	mid := left + (right-left)/2

	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
