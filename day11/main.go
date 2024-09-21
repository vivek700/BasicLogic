package main

import (
	"fmt"
)

func findDuplicates[T comparable](s []T) string {
	// for i, v := range s {
	// 	for j, value := range s {
	// 		if i == j {
	// 			continue
	// 		} else {
	// 			if v == value {
	// 				return "Duplicates found"
	// 			}
	// 		}
	// 	}
	// }
	m := make(map[T]int)
	for _, v := range s {
		_, ok := m[v]

		if ok {
			return "Duplicates found"
		} else {
			m[v] = 1
		}
	}
	return "No duplicates"

}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 1}
	numbers1 := []int{10, 20, 30, 40, 50}

	orders := []int{101, 102, 103, 104, 101}
	orders1 := []int{201, 202, 203, 204}
	sales := []int{5000, 6000, 7000, 5000, 8000}
	sales1 := []int{4000, 3000, 2000, 1000}

	cart := []string{"Rice", "Dal", "Oil", "Rice"}
	cart1 := []string{"Milk", "Bread", "Eggs"}

	fmt.Println(findDuplicates(numbers))
	fmt.Println(findDuplicates(numbers1))

	fmt.Println(findDuplicates(orders))
	fmt.Println(findDuplicates(orders1))

	fmt.Println(findDuplicates(sales))
	fmt.Println(findDuplicates(sales1))

	fmt.Println(findDuplicates(cart))
	fmt.Println(findDuplicates(cart1))

}
