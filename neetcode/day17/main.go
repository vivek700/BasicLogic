package main

import "fmt"

func characterReplacement(str string, k int) int {
	maxLen := 0
	maxValue := 0
	l := 0

	count := make(map[rune]int)
	for r, v := range str {
		_, ok := count[v]
		if !ok {
			count[v] = 0

		}
		count[v] = 1 + count[v]

		for _, value := range count {
			if value > maxValue {
				maxValue = value
			}
		}

		for (r-l+1)-maxValue > k {
			count[rune(str[l])]--
			l++
		}
		maxLen = max(maxLen, r-l+1)

	}

	return maxLen

}

func main() {
	s := "XYYX"
	s1 := "AAABABB"
	s2 := "XXXX"
	s3 := "ABBB"
	s4 := "AABABBA"
	s5 := "AABA"
	s6 := "ABAB"

	fmt.Println(characterReplacement(s, 2))
	fmt.Println(characterReplacement(s1, 1))
	fmt.Println(characterReplacement(s2, 1))
	fmt.Println(characterReplacement(s3, 2))
	fmt.Println(characterReplacement(s4, 1))
	fmt.Println(characterReplacement(s5, 0))
	fmt.Println(characterReplacement(s6, 0))

}
