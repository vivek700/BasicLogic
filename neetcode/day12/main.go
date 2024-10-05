package main

import "fmt"

func threeSum(n []int) [][]int {
	var newNum [][]int

	l, r := 0, len(n)-1

	for l < r {
		for k := range n {
			if k != l && k != r {
				if n[l]+n[r]+n[k] == 0 {
					newNum = append(newNum, []int{n[l], n[r], n[k]})
				} else {
					l++
					r--
				}
			}
		}
	}
	return newNum
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))

}
