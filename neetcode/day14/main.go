package main

import "fmt"

func trap(heights []int) int {
	l := 0
	trapWater := 0

	for l < len(heights) {
		if heights[l] > 0 {
			r := l + 1
			h := 0
			for r < len(heights) {
				if heights[l] > heights[r] {
					h += heights[r]
					r++
				} else if r-l == 1 {
					break

				} else {
					trapWater += heights[l]*(r-l-1) - h
					break
				}
			}
			l = r

		} else {
			l++
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
