package main

import "fmt"

func maxProfit(prices []int) int {
	maxPro := 0
	l := 0
	r := 1
	for l < r && r < len(prices) {
		if prices[r]-prices[l] < 0 {
			l = r
		} else {
			p := prices[r] - prices[l]
			if p > maxPro {
				maxPro = p
			}
		}
		r++
	}
	return maxPro
}

func main() {
	prices := []int{10, 1, 5, 6, 7, 1}
	prices1 := []int{10, 8, 7, 5, 2}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit(prices1))
}
