package main

import (
	"fmt"
)

func maxSlidingWindow(n []int, k int) []int {
	l, r := 0, 0

	res := []int{}
	q := []int{}

	for r < len(n) {

		for len(q) > 0 && n[q[len(q)-1]] < n[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if l > q[0] {
			q = q[1:]
		}

		if r+1 >= k {
			res = append(res, n[q[0]])
			l++
		}

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
