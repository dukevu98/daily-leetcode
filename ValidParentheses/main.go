package main

import "fmt"

func main() {
	//fmt.Println(isValid("(}"))
	//fmt.Println(isValid("({}[])"))
	//fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("()[]{}"))

}

var charMap = map[uint8]uint8{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}
	stack := []uint8{}
	for i := 0; i < len(s); i++ {
		if opening, ok := charMap[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
