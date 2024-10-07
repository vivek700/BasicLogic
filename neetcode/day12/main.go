package main

import (
	"fmt"
	"slices"
)

func threeSum(n []int) [][]int {
	var newNum [][]int

	l, r := 0, 1

	for r < len(n) {
		for k := range n {
			if k != l && k != r {
				if n[l]+n[r]+n[k] == 0 {
					newNum = append(newNum, []int{n[l], n[r], n[k]})

				}
			}
		}
		r++

	}
	for _, v := range newNum {
		slices.Sort(v)
	}

	return newNum
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums1 := []int{0, 1, 1}

	fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums1))

}
