package main

import (
	"fmt"
	"strings"
)

func main() {
	request1 := numberToWords(1234567892)
	resutl2 := numberToWords(20)
	fmt.Println("result")
	fmt.Println(request1)
	fmt.Println(resutl2 == "Twenty")
}

var numbers = map[int]string{
	0:  "Zero",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

func numberToWords(num int) string {
	var result string
	if num == 0 {
		return "Zero"
	}
	if num >= 1e9 {
		result += numberToWords(num/1e9) + " Billion "
		num %= 1e9
	}
	if num >= 1e6 {
		result += numberToWords(num/1e6) + " Million "
		num %= 1e6
	}
	if num >= 1e3 {
		result += numberToWords(num/1e3) + " Thousand "
		num %= 1e3
	}
	if num >= 1e2 {
		result += numberToWords(num/1e2) + " Hundred "
		num %= 1e2
	}
	if num >= 20 {
		result += numbers[num/10*10] + " "
		num %= 10
	}
	if num > 0 {
		result += numbers[num]
	}
	result = strings.TrimSpace(result)
	return result
}
