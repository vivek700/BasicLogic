package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}

}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)

	} else {
		this.minStack = append(this.minStack, min(val, this.minStack[len(this.minStack)-1]))
	}

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(5)
	fmt.Println(obj.stack)
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.stack)

}
