package main

import (
	"fmt"
)

func isAlphaNumeric(n byte) bool {
	if n != ' ' && 65 <= n && n <= 90 || n >= 97 && n <= 122 || '0' <= n && n <= '9' {
		return true

	}
	return false
}

func isSameChar(cl, cr int) bool {

	if cr < 65 || cl < 65 {
		if cl == cr {
			return true

		}
		return false

	}
	if cl-cr == 0 || cl-cr == 32 || cl-cr == -32 {
		return true
	}
	return false
}

func isPalindrome(s string) bool {

	leftptr := 0
	rightptr := len(s) - 1
	for leftptr != rightptr && leftptr < rightptr {
		fmt.Println(leftptr)
		fmt.Println(rightptr)
		if !isAlphaNumeric(s[leftptr]) {
			leftptr++
			continue
		}
		if !isAlphaNumeric(s[rightptr]) {
			rightptr--
			continue
		}
		if isSameChar(int(s[leftptr]), int(s[rightptr])) {
			leftptr++
			rightptr--
		} else {

			fmt.Println(s[leftptr], s[rightptr])
			return false
		}

	}
	return true

}

func main() {
	str := "Was it a car or a cat I saw?"
	s := "tab a cat"
	s1 := "0P"
	s2 := "A man, a plan, a canal: Panama"
	s3 := "1b1"

	fmt.Println(isPalindrome(str))
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
	fmt.Println(isPalindrome(s3))

}
