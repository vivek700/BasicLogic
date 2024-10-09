package main

import "fmt"

func twoSum(n []int, target int) []int {

	i, j := 0, len(n)-1

	for i < j {
		if n[i]+n[j] > target {
			j--

		} else if n[i]+n[j] < target {
			i++

		} else {
			return []int{i + 1, j + 1}
		}
	}

	return nil

}

func main() {

	nums := []int{1, 2, 3, 4}
	nums1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(twoSum(nums, 3))
	fmt.Println(twoSum(nums1, 10))

}
