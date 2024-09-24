package main

import "fmt"

func findSecondLargestNum(s []int) interface{} {
	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
	// return "Not enough unique elements"
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println(findSecondLargestNum(numbers))
}
