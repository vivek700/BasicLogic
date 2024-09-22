package main

import (
	"fmt"
	"slices"
)

func removeDuplicates[T comparable](s []T) []T {

	// m := make(map[T]int)

	// var slice []T
	// for _, v := range s {
	// 	_, ok := m[v]
	// 	if ok {
	// 		continue
	// 	} else {
	// 		m[v] = 1
	// 		slice = append(slice, v)
	// 	}
	// }
	var slice []T
	for _, v := range s {
		if !slices.Contains(slice, v) {
			slice = append(slice, v)

		}
	}

	return slice
}

func main() {
	numbers := []int{1, 2, 2, 3, 4, 4, 5}
	numbers1 := []int{7, 8, 9, 7, 8, 9, 10}

	names := []string{"Ankit", "Ravi", "Ankit", "Neha", "Ravi", "Priya"}
	names1 := []string{"Pooja", "Suresh", "Pooja", "Suresh", "Arjun"}

	product_ids := []int{101, 102, 103, 101, 104, 105, 102}
	product_ids1 := []int{201, 202, 203, 201, 204, 202}

	contacts := []int{9876543210, 9876543211, 9876543210, 9876543212, 9876543211}
	contacts1 := []int{9123456789, 9123456780, 9123456789, 9123456781}

	fmt.Println(removeDuplicates(numbers))
	fmt.Println(removeDuplicates(numbers1))

	fmt.Println(removeDuplicates(names))
	fmt.Println(removeDuplicates(names1))

	fmt.Println(removeDuplicates(product_ids))
	fmt.Println(removeDuplicates(product_ids1))

	fmt.Println(removeDuplicates(contacts))
	fmt.Println(removeDuplicates(contacts1))
}
