package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	var newstring []string
	for i := len(s) - 1; i >= 0; i-- {
		newstring = append(newstring, string(s[i]))

	}
	return strings.Join(newstring, "")

}

func reverseStringNotOrder(s string) string {

	spitSlice := strings.Split(s, " ")
	newRslice := []string{}
	for _, v := range spitSlice {
		newRslice = append(newRslice, reverseString(v))

	}

	return strings.Join(newRslice, " ")

}

func isPalindrome(s string) string {
	if s == reverseString(s) {
		return "Palindrome"
	}
	return "Not a palindrome"
}

func reverseOrder(s string) string {
	sp := strings.Split(s, " ")
	reverseOrderS := []string{}
	for i := len(sp) - 1; i >= 0; i-- {
		reverseOrderS = append(reverseOrderS, sp[i])

	}
	return strings.Join(reverseOrderS, " ")
}

func main() {
	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString("OpenAI"))
	fmt.Println(reverseStringNotOrder("Hello World"))
	fmt.Println(reverseStringNotOrder("OpenAI is cool"))

	fmt.Println(isPalindrome("OpenAI is cool"))
	fmt.Println(isPalindrome("madam"))
	fmt.Println(isPalindrome("hello"))
	fmt.Println(reverseOrder("Learning to code is fun"))
	fmt.Println(reverseOrder("I love programming"))

}
