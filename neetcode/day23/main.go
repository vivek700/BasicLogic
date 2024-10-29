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
	res = n.nums[len(n.nums)-1] + n.nums[len(n.nums)-2]
	n.nums = n.nums[:len(n.nums)-2]
	n.nums = append(n.nums, res)
}
func (n *Stack) Subtract() {
	res := 0
	res = n.nums[len(n.nums)-2] - n.nums[len(n.nums)-1]
	n.nums = n.nums[:len(n.nums)-2]
	n.nums = append(n.nums, res)
}
func (n *Stack) Mul() {
	res := 0
	res = n.nums[len(n.nums)-1] * n.nums[len(n.nums)-2]
	n.nums = n.nums[:len(n.nums)-2]
	n.nums = append(n.nums, res)
}
func (n *Stack) Divide() {
	res := 0
	res = n.nums[len(n.nums)-2] / n.nums[len(n.nums)-1]
	fmt.Println(res)
	n.nums = n.nums[:len(n.nums)-2]
	fmt.Println(n.nums)
	n.nums = append(n.nums, res)
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

	tokens := []string{"1", "2", "+", "3", "*", "4", "-"}
	tokens1 := []string{"4", "13", "5", "/", "+"}

	fmt.Println(evalRPN(tokens))
	fmt.Println(evalRPN(tokens1))

}
