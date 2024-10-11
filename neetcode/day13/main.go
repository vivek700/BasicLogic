package main

import "fmt"

func maxArea(nums []int) int {
	l, r := 0, len(nums)-1

	maxAmount := 0
	for l < r {
		if nums[l] < nums[r] {
			water := nums[l] * (r - l)
			if maxAmount < water {
				maxAmount = water
			}
			l++

		} else {
			water := nums[r] * (r - l)
			if maxAmount < water {
				maxAmount = water
			}
			r--

		}
	}
	return maxAmount
}

func main() {
	nums := []int{1, 7, 2, 5, 4, 7, 3, 6}
	nums1 := []int{2, 2, 2}

	fmt.Println(maxArea(nums))
	fmt.Println(maxArea(nums1))
}
