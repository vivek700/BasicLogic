package main

import (
	"fmt"
	"slices"
)

func threeSum(n []int) [][]int {

	var newNum [][]int

	slices.Sort(n)

	for i, k := range n {
		if i > 0 && k == n[i-1] {
			continue
		}
		l, r := i+1, len(n)-1
		for l < r {
			sum := k + n[l] + n[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++

			} else {
				newNum = append(newNum, []int{k, n[l], n[r]})
				l++
				for l < r {
					if n[l] == n[l-1] {
						l++
					} else {
						break
					}
				}

			}

		}

	}

	return newNum
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums1 := []int{0, 1, 1}
	nums2 := []int{0, 0, 0}

	fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums1))
	fmt.Println(threeSum(nums2))

}
