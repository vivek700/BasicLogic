package main

import "fmt"

func sumNumbers(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i

	}
	return sum
}

func main() {

	fmt.Println(sumNumbers(10))

}
