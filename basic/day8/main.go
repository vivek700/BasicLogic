package main

import (
	"fmt"
)

func findMaximum(s []int) int {
	// return slices.Max(s)

	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max

}

func main() {

	numbers := []int{12, 45, 13, 67, 34}
	numbers1 := []int{5, 1, 7, 6, 9}
	fmt.Println(findMaximum(numbers))
	fmt.Println(findMaximum(numbers1))

}
