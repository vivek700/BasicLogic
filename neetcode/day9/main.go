package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {

	res := map[int]int{}

	for _, v := range nums {
		if _, ok := res[v]; !ok {
			res[v] = v

		}

	}
	longestseq := 0
	resmap := make(map[int]int)
	for _, v := range res {
		if _, ok := res[v-1]; !ok {
			length := 0
			for {
				if _, ok := res[v+length]; !ok {
					break
				}
				length++

			}
			resmap[length] = length
		}
	}

	for key := range resmap {
		if key > longestseq {
			longestseq = key
		}
	}

	return longestseq
}

func main() {
	input := []int{2, 20, 4, 10, 3, 4, 5}
	input1 := []int{0, 3, 2, 5, 4, 6, 1, 1}
	input2 := []int{0}
	input3 := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}

	fmt.Println(longestConsecutive(input))
	fmt.Println(longestConsecutive(input1))
	fmt.Println(longestConsecutive(input2))
	fmt.Println(longestConsecutive(input3))

}
