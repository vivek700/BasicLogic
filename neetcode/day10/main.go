package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	newStr := ""
	newReverseStr := ""

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " && 65 <= s[i] && s[i] <= 90 || s[i] >= 97 && s[i] <= 122 || '0' <= s[i] && s[i] <= '9' {
			newReverseStr += strings.ToLower(string(s[i]))
		}
		if string(s[(len(s)-1)-i]) != " " && 65 <= s[(len(s)-1)-i] && s[(len(s)-1)-i] <= 90 || s[(len(s)-1)-i] >= 97 && s[(len(s)-1)-i] <= 122 || '0' <= s[(len(s)-1)-i] && s[(len(s)-1)-i] <= '9' {
			newStr += strings.ToLower(string(s[(len(s)-1)-i]))
		}

	}

	if newReverseStr == newStr {
		return true
	}

	return false

}

func main() {
	str := "Was it a car or a cat I saw?"
	s := "tab a cat"
	s1 := "0P"

	fmt.Println(isPalindrome(str))
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(s1))

}
