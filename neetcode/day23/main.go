package main

import (
	"fmt"
	"log"
	"strconv"
)

type Stack struct {
	nums []int
}

func (n *Stack) Add() {
	res := 0
	for _, v := range n.nums {
		res += v
	}
	n.nums = []int{res}
}
func (n *Stack) Subtract() {
	res := n.nums[0]
	for _, v := range n.nums {
		res -= v
	}
	n.nums = []int{res}
}
func (n *Stack) Divide() {
	res := n.nums[0]
	for _, v := range n.nums {
		res /= v
	}
	n.nums = []int{res}
}
func (n *Stack) Mul() {
	res := n.nums[0]
	for _, v := range n.nums {
		res *= v
	}
	n.nums = []int{res}
}

func evalRPN(tokens []string) int {

	stack := Stack{}

	for _, v := range tokens {
		if v == "+" {
			stack.Add()
		} else if v == "-" {
			stack.Subtract()

		} else if v == "*" {
			stack.Mul()
		} else if v == "/" {
			stack.Divide()
		} else {
			intNum, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			stack.nums = append(stack.nums, intNum)
		}
	}
	return stack.nums[0]

}

func main() {

	tokens := []string{"1", "2", "+", "3", "*"}

	fmt.Println(evalRPN(tokens))

}
