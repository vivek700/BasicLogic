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
			if sum2 == sumOfS1 {
				return true
			}
		} else {
			l = r - 1
			r = l - 1
			sum2 = 0
		}
		r++
	}

	return false
}

func main() {

	s1, s2 := "abc", "lecabee"
	s3, s4 := "abc", "lecaabee"
	fmt.Println(checkInclusion(s1, s2))
	fmt.Println(checkInclusion(s3, s4))

}
