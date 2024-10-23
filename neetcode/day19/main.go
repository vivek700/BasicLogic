package main

import (
	"fmt"
	"math"
)

func minWindow(s, t string) string {
	if t == "" {
		return ""
	}

	tMap := make(map[rune]int)
	for _, v := range t {
		if _, ok := tMap[v]; !ok {
			tMap[v] = 1
		} else {
			tMap[v]++
		}

	}

	l := 0

	sMap := make(map[rune]int)
	res := s[0:]
	reslen := float64(math.Inf(+1))
	have, need := 0, len(tMap)

	for r, value := range s {
		if _, sOk := sMap[value]; sOk {
			sMap[value]++
		} else {
			sMap[value] = 1
		}
		if _, ok := tMap[value]; ok && sMap[value] == tMap[value] {
			have++

		}
		for have == need {
			if float64(r-l+1) < reslen {
				res = s[l : r+1]
				reslen = float64(r - l + 1)
			}
			sMap[rune(s[l])]--
			if _, has := tMap[rune(s[l])]; has && sMap[rune(s[l])] < tMap[rune(s[l])] {
				have--

			}
			l++
		}

	}
	if reslen != float64(math.Inf(1)) {
		return res
	} else {
		return ""
	}

}

func main() {

	s, t := "OUZODYXAZV", "XYZ"
	s1, t1 := "a", "b"
	s2, t2 := "a", "a"

	fmt.Println(minWindow(s, t))
	fmt.Println(minWindow(s1, t1))
	fmt.Println(minWindow(s2, t2))

}
