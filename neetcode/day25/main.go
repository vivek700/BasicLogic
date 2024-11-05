package main

import (
	"fmt"
)

type Stack struct {
	item [][]int
}

func (s *Stack) Pop() (int, int) {

	temp := s.item[len(s.item)-1][0]
	index := s.item[len(s.item)-1][1]
	s.item = s.item[:len(s.item)-1]
	return temp, index

}

func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))
	stack := &Stack{item: make([][]int, 0)}

	for i, v := range temperatures {
		for len(stack.item) != 0 && stack.item[len(stack.item)-1][0] < v {
			_, ti := stack.Pop()
			res[ti] = (i - ti)
		}
		stack.item = append(stack.item, []int{v, i})
	}

	return res

}

func main() {
	temps := []int{30, 38, 30, 36, 35, 40, 28}
	fmt.Println(dailyTemperatures(temps))

}
