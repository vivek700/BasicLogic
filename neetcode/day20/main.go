package main

import (
	"fmt"
	"math"
)

func maxSlidingWindow(n []int, k int) []int {

	l, r := 0, k

	res := []int{}

	for r < len(n)+1 {
		max := math.Inf(-1)
		for _, v := range n[l:r] {
			if float64(v) > max {
				max = float64(v)
			}

		}
		res = append(res, int(max))

		l++
		r++
	}

	return res
}

func main() {

	nums := []int{1, 2, 1, 0, 4, 2, 6}
	nums1 := []int{1, -1}

	fmt.Println(maxSlidingWindow(nums, 3))
	fmt.Println(maxSlidingWindow(nums1, 1))

}
