package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	var result [][]int
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r, c := rStart, cStart
	step := 1

	result = append(result, []int{r, c})
	for len(result) < rows*cols {
		for i := 0; i < len(directions); i++ {
			direction := directions[i]
			for j := 0; j < step; j++ {
				r += direction[0]
				c += direction[1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					result = append(result, []int{r, c})
				}
			}
			if i%2 == 1 {
				step++
			}
		}
	}
	return result
}
