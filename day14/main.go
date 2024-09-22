package main

import "fmt"

func findDifference(s []int, threshold int) string {
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] > threshold {
			return "Stock is going up"

		}

	}
	return "Stock is stable"
}

func main() {
	prices := []int{100, 105, 102, 110, 115}
	prices1 := []int{100, 101, 102, 103, 105}
	prices2 := []int{200, 210, 215, 220, 225}
	prices3 := []int{200, 210, 220, 230, 240}
	prices4 := []int{500, 505, 510, 515, 520}
	prices5 := []int{500, 510, 520, 530, 540}
	prices6 := []int{1200, 1210, 1220, 1230, 1240}
	prices7 := []int{1200, 1220, 1240, 1260, 1280}

	fmt.Println(findDifference(prices, 5))
	fmt.Println(findDifference(prices1, 5))
	fmt.Println(findDifference(prices2, 8))
	fmt.Println(findDifference(prices3, 8))
	fmt.Println(findDifference(prices4, 6))
	fmt.Println(findDifference(prices5, 6))
	fmt.Println(findDifference(prices6, 12))
	fmt.Println(findDifference(prices7, 12))
}
