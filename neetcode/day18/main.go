package main

import (
	"fmt"
)

func checkInclusion(s1, s2 string) bool {

	s1Array := [26]int{}
	for _, v := range s1 {
		index := v - 'a'
		s1Array[index]++

	}

	l, r := 0, len(s1)
	for r < len(s2)+1 {
		s2Array := [26]int{}
		for _, v := range s2[l:r] {
			index := v - 'a'
			s2Array[index]++

		}

		if s1Array == s2Array {
			return true
		}
		l++
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
