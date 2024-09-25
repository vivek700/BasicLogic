package main

import "fmt"

func findLargestAndSmallestNumber(s []int) (int, int) {
	smallest := s[0]
	largest := s[0]

	for _, v := range s {
		if v > largest {
			largest = v
		} else if v < smallest {
			smallest = v
		}
	}
	return largest, smallest

}

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	numbers1 := []int{-10, -20, -30, -5, -1}
	temp := []int{32, 35, 28, 30, 33, 31, 29}
	temp1 := []int{15, 12, 18, 10, 20, 11, 14}
	rainfall := []int{5, 12, 8, 20, 10, 15, 25}
	rainfall1 := []int{2, 0, 0, 4, 5, 1, 3}
	scores := []int{85, 92, 88, 79, 95, 91, 82}
	scores1 := []int{45, 56, 74, 62, 68, 80, 55}

	fmt.Println(findLargestAndSmallestNumber(numbers))
	fmt.Println(findLargestAndSmallestNumber(numbers1))
	fmt.Println(findLargestAndSmallestNumber(temp))
	fmt.Println(findLargestAndSmallestNumber(temp1))
	fmt.Println(findLargestAndSmallestNumber(rainfall))
	fmt.Println(findLargestAndSmallestNumber(rainfall1))
	fmt.Println(findLargestAndSmallestNumber(scores))
	fmt.Println(findLargestAndSmallestNumber(scores1))

}
