package main

import (
	"fmt"
)

func trap(height []int) int {
	l, r := 0, len(height)-1
	trapWater := 0
	h := 0
	for l < r {
		if height[l] == 0 {
			l++
		}
		if height[r] == 0 {
			r--
		}
		if height[l] < height[r] {
			trapWater += height[l] * (r - l - 1)
			l++
		} else {
			trapWater += height[r] * (r - l - 1)
			r--

		}

	}
	return trapWater
}

func main() {
	heights := []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}
	heights1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(heights))
	fmt.Println(trap(heights1))

}
