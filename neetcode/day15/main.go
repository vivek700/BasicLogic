package main

import "fmt"

func maxProfit(prices []int) int {
	maxPro := 0
	for i, v := range prices {
		for j := i + 1; j < len(prices); j++ {
			calP := prices[j] - v
			if calP > maxPro {
				maxPro = calP

			}

		}
	}
	return maxPro
}

func main() {
	prices := []int{10, 1, 5, 6, 7, 1}
	prices1 := []int{10, 8, 7, 5, 2}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit(prices1))
}
