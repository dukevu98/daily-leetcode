package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	var l = x
	var r int
	for x > 0 {
		r = r*10 + x%10
		x = x / 10
	}
	return l == r
}
