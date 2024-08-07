package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var RomanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		result += RomanMap[string(s[i])]
		if i > 0 && RomanMap[string(s[i-1])] < RomanMap[string(s[i])] {
			result -= 2 * RomanMap[string(s[i-1])]
		}
	}
	return result
}
