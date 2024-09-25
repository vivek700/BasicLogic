package main

import "fmt"

func findDuplicates(s []int) bool {
	elements := make(map[int]int)

	for _, v := range s {
		if _, ok := elements[v]; ok {
			return true
		} else {
			elements[v] = 1
		}
	}
	return false
}

func main() {

	numbers := []int{1, 2, 3, 3}
	numbers1 := []int{1, 2, 3, 4}

	fmt.Println(findDuplicates(numbers))
	fmt.Println(findDuplicates(numbers1))

}
