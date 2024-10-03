package main

import "fmt"

func twoSum(n []int, target int) []int {
	diff := 0
	for i, v := range n {
		diff = target - v
		for j, v1 := range n {
			if i == j {
				continue

			}
			if diff == v1 {
				return []int{i + 1, j + 1}
			}
		}

	}
	return nil

}

func main() {

	nums := []int{1, 2, 3, 4}
	fmt.Println(twoSum(nums, 3))

}
