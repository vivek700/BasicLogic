package main

import "fmt"

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{}

}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	min := this.stack[0]
	for _, v := range this.stack {
		if v < min {
			min = v
		}
	}
	return min

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
