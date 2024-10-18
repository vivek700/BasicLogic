package main

import (
	"fmt"
)

func checkInclusion(s1, s2 string) bool {
	sumOfS1 := 0
	sum2 := 0
	for _, v := range s1 {
		sumOfS1 += int(v)

	}
	l, r := 0, 0
	for r < len(s2) {

		if r-l < len(s1) {
			sum2 += int(s2[r])
			if sum2 == sumOfS1 && (s2[l] == s1[0] || s2[l] == s1[len(s1)-1] || s2[r] == s1[0] || s2[r] == s1[len(s1)-1]) {
				return true
			}
		} else {
			l++
			r--
			sum2 -= int(s2[l-1])
		}
		r++
	}

	return false
}

func main() {

	s1, s2 := "abc", "lecabee"
	s3, s4 := "abc", "lecaabee"
	s5 := "ab"
	s6 := "eidbaooo"
	s7 := "adc"
	s8 := "dcda"
	s9 := "abc"
	s10 := "ccccbbbbaaaa"
	s11 := "abc"
	s12 := "bbbca"
	fmt.Println(checkInclusion(s1, s2))
	fmt.Println(checkInclusion(s3, s4))
	fmt.Println(checkInclusion(s5, s6))
	fmt.Println(checkInclusion(s7, s8))
	fmt.Println(checkInclusion(s9, s10))
	fmt.Println(checkInclusion(s11, s12))

}
