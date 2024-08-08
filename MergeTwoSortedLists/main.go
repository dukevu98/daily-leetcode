package main

import (
	"fmt"
	"leetcode/Utils"
)

func main() {
	list1 := Utils.BuildList([]int{1, 2, 4})
	list2 := Utils.BuildList([]int{1, 3, 4})
	result := mergeTwoLists(list1, list2)
	fmt.Println(result)
}

func mergeTwoLists(list1 *Utils.ListNode, list2 *Utils.ListNode) *Utils.ListNode {
	result := &Utils.ListNode{}
	current := result
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return result.Next
}
