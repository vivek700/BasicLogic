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

func toLowerCase(ch byte) byte {

	if ch >= 65 && ch <= 90 {
		return ch + 32
	}
	return ch

}

func isPalindrome(s string) bool {

	leftptr := 0
	rightptr := len(s) - 1
	for leftptr != rightptr && leftptr < rightptr {

		if !isAlphaNumeric(s[leftptr]) {
			leftptr++
			continue
		}
		if !isAlphaNumeric(s[rightptr]) {
			rightptr--
			continue
		}
		if toLowerCase(s[leftptr]) == toLowerCase(s[rightptr]) {
			leftptr++
			rightptr--
		} else {
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
