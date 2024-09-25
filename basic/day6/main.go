package main

import "fmt"

func sumNumbers(n int) int {
	// var sum int = 0
	// for i := 1; i <= n; i++ {
	// 	sum += i

	// }
	// return sum
	return (n * (n + 1)) / 2
}

func calculateCandles(n int) int {
	return n * (n + 1) / 2
}

func calculateTotalSavings(days, daily_savings int) int {
	return days * daily_savings
}

var calculateSales = func(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum

}

func main() {

	fmt.Println(sumNumbers(10))
	fmt.Println(calculateTotalSavings(7, 10))
	fmt.Println(calculateTotalSavings(5, 20))
	fmt.Println(calculateCandles(4))
	fmt.Println(calculateCandles(6))

	sales := []float64{50.5, 60.0, 45.75, 70.25, 55.0, 80.0, 65.5}

	fmt.Printf("%.2f\n", calculateSales(sales))

	sales1 := []float64{20.0, 30.0, 25.0, 40.0}

	fmt.Printf("%.2f\n", calculateSales(sales1))

}
