package main

import "fmt"

func findCommonNumbers(s, y []int) []int {
	var array []int
	for _, value := range s {
		for _, v := range y {
			if v == value {
				array = append(array, v)
			}

		}
	}
	return array
}

func findCommonNames(s, y []string) []string {
	var slice []string

	for _, value := range s {
		for _, v := range y {
			if value == v {
				slice = append(slice, v)
			}
		}
	}
	return slice
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	fmt.Println(findCommonNumbers(s1, s2))

	s3 := []int{1, 2, 3}
	s4 := []int{4, 5, 6}
	fmt.Println(findCommonNumbers(s3, s4))

	list1 := []string{"Alice", "Bob", "Charlie"}
	list2 := []string{"Bob", "David", "Alice"}

	fmt.Println(findCommonNames(list1, list2))

	list3 := []string{"Alice", "Bob"}
	list4 := []string{"Charlie", "David"}

	fmt.Println(findCommonNames(list3, list4))

	list5 := []string{"Math", "Science", "History"}
	list6 := []string{"Science", "Art", "History"}

	fmt.Println(findCommonNames(list5, list6))

	recipe1 := []string{"Flour", "Sugar", "Eggs"}
	recipe2 := []string{"Sugar", "Butter", "Eggs"}

	fmt.Println(findCommonNames(recipe1, recipe2))

}
