package main

import (
	"fmt"
)

func trap(height []int) int {
	trapWater := 0
	l, r := 0, len(height)-1
	maxL, maxR := 0, height[r]

	for l < r {
		if maxL < maxR {
			waterDrop := maxL - height[l]
			if waterDrop > 0 {
				trapWater += waterDrop
				maxL = height[l]

			}
			l++

		} else {
			waterDrop := maxR - height[l]
			if waterDrop > 0 {
				trapWater += waterDrop
				maxR = height[r]

			}
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
