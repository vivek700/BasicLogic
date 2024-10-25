package main

import "fmt"

func isValid(s string) bool {

	stack := []string{}
	close := map[string]string{")": "(", "]": "[", "}": "{"}

	for _, v := range s {
		if _, ok := close[string(v)]; ok {
			if len(stack) != 0 && stack[len(stack)-1] == close[string(v)] {
				stack = stack[:len(stack)-1]
			} else {
				return false

			}

		} else {
			stack = append(stack, string(v))
		}

	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}

func main() {

	s := "[]"
	s1 := "([{}])"
	s2 := "[(])"
	s3 := "()[]{}"

	fmt.Println(isValid(s))
	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
}
