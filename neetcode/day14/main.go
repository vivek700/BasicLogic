package main

import (
	"fmt"
)

func trap(height []int) int {
	l, r := 0, len(height)-1
	trapWater := 0
	for l < r {
		if height[l] == 0 {
			l++
		} else {
			k := l + 1
			for k < r {
				if height[k] < height[l] && height[k] < height[r] {
					if height[l] < height[r] {
						trapWater += height[l] * (r - l - 1)

					} else {
						trapWater += height[r] * (r - l - 1)

					}

				} else {
					break
				}
			}

			if height[l] > height[r] {
				r--
			} else {
				l++
			}
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
