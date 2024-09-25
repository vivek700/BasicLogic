package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, b+a
		return ret
	}
}

func nthNumberOfFib(n int) int {
	f := fibonacci()
	nth := 0
	for i := 0; i < n; i++ {
		nth = f()
	}
	return nth
}
func findfiboNumbers(n int) []int {
	var slice []int
	f := fibonacci()
	for i := 0; i < n; i++ {
		slice = append(slice, f())

	}
	return slice

}

func closestFiboNum(budget int) int {
	ret := 0
	f := fibonacci()
	for {
		if fnum := f(); fnum <= budget {
			ret = fnum
		} else {
			break
		}
	}
	return ret

}

func main() {
	fmt.Println(nthNumberOfFib(5))
	fmt.Println(nthNumberOfFib(7))

	fmt.Println(findfiboNumbers(6))
	fmt.Println(closestFiboNum(50))
	fmt.Println(closestFiboNum(15))
	fmt.Println(findfiboNumbers(5))
	fmt.Println(findfiboNumbers(7))
	fmt.Println(findfiboNumbers(15))
	fmt.Println(closestFiboNum(3))
	fmt.Println(closestFiboNum(90))
	fmt.Println(closestFiboNum(87))
	fmt.Println(closestFiboNum(89))
	fmt.Println(closestFiboNum(7))

}
