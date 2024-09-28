package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {

	newNums := make([][]int, len(nums)+1)
	freqElement := make(map[int]int)

	for _, v := range nums {
		if _, ok := freqElement[v]; !ok {
			freqElement[v] = 1
		} else {
			freqElement[v] += 1

		}
	}
	for i, v := range freqElement {
		newNums[v] = append(newNums[v], i)
	}

	ret := []int{}
	for i := len(newNums) - 1; i > 0; i-- {
		for _, v := range newNums[i] {
			ret = append(ret, v)
			if len(ret) == k {
				return ret
			}
		}

	}
	return nil

}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	nums1 := []int{7, 7}
	nums2 := []int{3, 0, 1, 0}
	fmt.Println(topKFrequent(nums, 2))
	fmt.Println(topKFrequent(nums1, 1))
	fmt.Println(topKFrequent(nums2, 1))
}
