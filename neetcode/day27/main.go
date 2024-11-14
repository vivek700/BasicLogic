package main

import "fmt"

type Pair struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []Pair{}

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			index := stack[len(stack)-1].index
			height := stack[len(stack)-1].height
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, height*(i-index))
			start = index

		}
		stack = append(stack, Pair{index: start, height: h})
	}
	for _, p := range stack {
		maxArea = max(maxArea, p.height*(len(heights)-p.index))
	}
	return maxArea

}

func main() {

	heights := []int{7, 1, 7, 2, 2, 4}

	fmt.Println(largestRectangleArea(heights))
}
