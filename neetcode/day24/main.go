package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	res := []string{}
	stack := []string{}

	var backTrack func(openN, closeN int)

	backTrack = func(openN, closeN int) {
		if openN == n && closeN == n {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if openN < n {
			stack = append(stack, "(")
			backTrack(openN+1, closeN)
			stack = stack[:len(stack)-1]
		}
		if closeN < openN {
			stack = append(stack, ")")
			backTrack(openN, closeN+1)
			stack = stack[:len(stack)-1]

		}

	}
	backTrack(0, 0)
	return res

}

func main() {

	n := 1
	n1 := 3

	fmt.Println(generateParenthesis(n))
	fmt.Println(generateParenthesis(n1))

}
