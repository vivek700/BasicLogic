package main

import "fmt"

func sumOfInteger(nums []int, target int) []int {

	// var indices []int
	// for i, v := range nums {
	// 	for j, value := range nums[i+1:] {
	// 		if v+value == target {
	// 			indices = append(indices, i, j+1+i)
	// 			return indices
	// 		}

	// 	}
	// }
	// return indices

	indices := make(map[int]int)
	for i, value := range nums {
		if v, ok := indices[target-value]; !ok {
			indices[value] = i
		} else {
			return []int{v, i}
		}
	}
	return []int{}

}

func main() {

	nums := []int{3, 4, 5, 6}
	nums1 := []int{4, 5, 6}
	nums2 := []int{5, 5}
	fmt.Println(sumOfInteger(nums, 7))
	fmt.Println(sumOfInteger(nums1, 10))
	fmt.Println(sumOfInteger(nums2, 10))

}
