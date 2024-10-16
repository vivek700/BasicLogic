package main

import (
	"fmt"
)

func lengthOfLongestSubstring(str string) int {
	// if len(str) == 1 {
	// 	return 1
	// }
	maxLen := 0
	start := 0
	charIndex := make(map[rune]int)

	for index, char := range str {
		if value, ok := charIndex[char]; ok && value >= start {
			start = charIndex[char] + 1
		} else {
			maxLen = max(maxLen, index-start+1)
		}
		charIndex[char] = index

	}

	return maxLen

}

func main() {
	s := "zxyzxyz"
	s1 := "xxxx"
	s3 := "pwwkew"
	s4 := "abcabcbb"
	s5 := "dvdf"
	s6 := " "
	s7 := "au"
	s8 := ""

	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
	fmt.Println(lengthOfLongestSubstring(s5))
	fmt.Println(lengthOfLongestSubstring(s6))
	fmt.Println(lengthOfLongestSubstring(s7))
	fmt.Println(lengthOfLongestSubstring(s8))

}
