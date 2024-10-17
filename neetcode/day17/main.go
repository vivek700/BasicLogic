package main

import "fmt"

func characterReplacement(str string, k int) int {
	maxLen := 0

	l, r := 0, 1

	for r < len(str) {
		if str[l] != str[r] {
			if k > 0 {
				k--
			} else {
				l = r
			}

		}
		maxLen = max(maxLen, r-l+1)
		r++

	}
	return maxLen

}

func main() {
	s := "XYYX"
	s1 := "AAABABB"
	s2 := "XXXX"
	s3 := "ABBB"

	fmt.Println(characterReplacement(s, 2))
	fmt.Println(characterReplacement(s1, 1))
	fmt.Println(characterReplacement(s2, 1))
	fmt.Println(characterReplacement(s3, 2))

}
