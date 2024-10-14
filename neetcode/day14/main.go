package main

import (
	"fmt"
)

func trap(height []int) int {
	trapWater := 0
	l, r := 0, len(height)-1
	maxL, maxR := height[l], height[r]

	for l < r {
		if maxL < maxR {
			l++
			if height[l] > maxL {
				maxL = height[l]
			}
			waterDrop := maxL - height[l]
			if waterDrop > 0 {
				trapWater += waterDrop

			}

		} else {
			r--
			if height[r] > maxR {
				maxR = height[r]
			}
			waterDrop := maxR - height[r]
			if waterDrop > 0 {
				trapWater += waterDrop

			}

		}

	}
	return trapWater

}

func main() {
	heights := []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}
	heights1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	heights2 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(heights))
	fmt.Println(trap(heights1))
	fmt.Println(trap(heights2))

}
