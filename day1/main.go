package main

import (
	"fmt"
)

func isNumberEven(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func fairShare(n int) string {
	if n%2 == 0 {
		return "Yes"
	}
	return "No"
}

func listOfNumbers(num []int) []string {
	slice := make([]string, len(num))

	for i, v := range num {
		if v%2 == 0 {
			slice[i] = "Even"
		} else {
			slice[i] = "Odd"
		}
	}
	return slice
}

func main() {

	fmt.Println(isNumberEven(2))
	fmt.Println(isNumberEven(4))
	fmt.Println(isNumberEven(5))

	fmt.Println(fairShare(10))
	fmt.Println(fairShare(7))

	list := []int{4, 7, 10}
	fmt.Println(listOfNumbers(list))

}
