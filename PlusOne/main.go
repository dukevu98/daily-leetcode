package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 9, 9}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))

}

func plusOne(digits []int) []int {
	debt := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] = digits[i] + 1
			if digits[i] > 9 {
				digits[i] = 0
				debt = 1
			}
		} else if debt > 0 {
			digits[i] = digits[i] + debt
			if digits[i] > 9 {
				digits[i] = 0
				debt = 1
			} else {
				debt = 0
			}
		}
		if i == 0 && debt > 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
