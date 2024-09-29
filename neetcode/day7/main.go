package main

import (
	"fmt"
)

func productExceptSelf(n []int) []int {

	result := make([]int, len(n))
	for i := range result {
		result[i] = 1
	}
	preFix := 1
	postFix := 1

	for i, v := range n {
		result[i] *= preFix
		preFix *= v

	}
	for i := range n {
		result[len(n)-i-1] *= postFix
		postFix *= n[len(n)-1-i]

	}

	return result

}

func main() {
	nums := []int{1, 2, 4, 6}
	nums1 := []int{-1, 0, 1, 2, 3}

	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelf(nums1))

}
