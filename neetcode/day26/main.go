package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Pair struct {
	p int
	s int
}

func carFleet(target int, position, speed []int) int {

	pair := []Pair{}

	for i, v := range position {
		pair = append(pair, Pair{v, speed[i]})
	}

	slices.SortFunc(pair, func(a, b Pair) int {
		return cmp.Compare(b.p, a.p)
	})

	stack := []float64{}

	for _, p := range pair {
		time := float64(target-p.p) / float64(p.s)
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)

}

func main() {

	fmt.Println(carFleet(10, []int{1, 4}, []int{3, 2}))
	fmt.Println(carFleet(10, []int{4, 1, 0, 7}, []int{2, 2, 1, 1}))

}
