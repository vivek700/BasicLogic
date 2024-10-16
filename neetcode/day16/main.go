package main

import (
	"fmt"
)

func lengthOfLongestSubstring(str string) int {
	maxLen := 1
	len := 0
	for i, v := range str {
		if i == 0 {
			continue
		}
		if int(str[i-1])+1 == int(v) {
			len++
		} else {
			if len > maxLen {
				maxLen += len
			}
			len = 0
		}
	}
	return maxLen
}

func main() {
	s := "zxyzxyz"
	s1 := "xxxx"

	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring(s1))

}
