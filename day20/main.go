package main

import "fmt"

func findSecondLargestNum(s []int) interface{} {
	max := 0
	secondMax := 0
	for _, v := range s {
		if v > max {
			secondMax = max
			max = v
		} else if v > secondMax && secondMax != max && v != max {
			secondMax = v

		}
	}
	if secondMax == 0 {
		return "Not enough unique elements"
	}
	return secondMax
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers1 := []int{8, 3, 6, 3, 2}
	numbers2 := []int{5, 5, 5, 5}
	scores := []int{85, 92, 78, 92, 88}
	scores1 := []int{100, 100, 100}
	heights := []int{150, 200, 250, 200, 300}
	heights1 := []int{120, 120, 120}
	times := []int{12, 15, 10, 10}
	times1 := []int{8, 8, 8}

	fmt.Println(findSecondLargestNum(numbers))
	fmt.Println(findSecondLargestNum(numbers1))
	fmt.Println(findSecondLargestNum(numbers2))
	fmt.Println(findSecondLargestNum(scores))
	fmt.Println(findSecondLargestNum(scores1))
	fmt.Println(findSecondLargestNum(heights))
	fmt.Println(findSecondLargestNum(heights1))
	fmt.Println(findSecondLargestNum(times))
	fmt.Println(findSecondLargestNum(times1))
}
