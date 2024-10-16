package main

import (
	"fmt"
	"slices"
)

func lengthOfLongestSubstring(str string) int {
	maxLen := 0
	l := 0
	charset := []string{}

	for r, v := range str {
		k := 0
		for k < len(charset) {
			if charset[k] != string(v) {
				k++
			} else {
				charset = slices.Delete(charset, l, l+1)
				l++
				break
			}
		}
		charset = append(charset, string(v))

		maxLen = max(maxLen, r-l+1)
	}
	// fmt.Println(maxLen)
	// return len(charset)
	fmt.Println(charset)
	return maxLen

}

func main() {
	s := "zxyzxyz"
	// s1 := "xxxx"
	s3 := "pwwkew"
	s4 := "abcabcbb"
	s5 := "dvdf"

	fmt.Println(lengthOfLongestSubstring(s))
	// fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
	fmt.Println(lengthOfLongestSubstring(s5))

}
