package main

import "fmt"

func findFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * findFactorial(n-1)
}

func findFactorialOfSales(s []int) []int {
	slice := make([]int, len(s))

	for i, v := range s {
		slice[i] = findFactorial(v)
	}
	return slice
}

func main() {

	fmt.Println(findFactorial(4))
	fmt.Println(findFactorial(5))
	fmt.Println(findFactorial(1))
	fmt.Println(findFactorial(0))
	fmt.Println(findFactorial(20))

	sales := []int{2, 3, 4}
	sales1 := []int{1, 5}

	fmt.Println(findFactorialOfSales(sales))
	fmt.Println(findFactorialOfSales(sales1))

	fmt.Println(findFactorial(3))
	fmt.Println(findFactorial(6))
}
