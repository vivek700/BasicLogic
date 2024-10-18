package main

import "fmt"

func characterReplacement(str string, k int) int {
	maxLen := 0
	counter := 0
	l := 0

	count := make(map[rune]int)
	for r, v := range str {
		count[v] = 1 + count[v]
		if counter < count[rune(str[l])] {
			counter = count[rune(str[l])]
		} else if counter < count[rune(str[r])] {
			counter = count[rune(str[r])]
		}
		if r-l+1-counter > k {
			counter--
			l++
		}
		maxLen = max(len(count), r-l+1)

	}
	fmt.Println(count)

	return maxLen

}

func main() {
	s := "XYYX"
	s1 := "AAABABB"
	s2 := "XXXX"
	s3 := "ABBB"
	s4 := "AABABBA"
	s5 := "AABA"

	fmt.Println(characterReplacement(s, 2))
	fmt.Println(characterReplacement(s1, 1))
	fmt.Println(characterReplacement(s2, 1))
	fmt.Println(characterReplacement(s3, 2))
	fmt.Println(characterReplacement(s4, 1))
	fmt.Println(characterReplacement(s5, 0))

}
