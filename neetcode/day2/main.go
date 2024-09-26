package main

import (
	"fmt"
	"slices"
	"strings"
)

func sort(s string) string {
	newS := strings.Split(s, "")
	slices.Sort(newS)
	return strings.Join(newS, "")
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if sort(s) != sort(t) {
		return false
	}

	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for i := range s {
		if _, ok := map1[string(s[i])]; ok {
			map1[string(s[i])] += 1
		} else {
			map1[string(s[i])] = 1
		}
		if _, ok := map2[string(t[i])]; ok {
			map2[string(t[i])] += 1
		} else {
			map2[string(t[i])] = 1
		}

	}
	for i, v := range map1 {
		if value, ok := map2[i]; !ok {
			return false
		} else if value != v {
			return false

		}
	}
	return true

}

func main() {

	fmt.Println(isAnagram("racecar", "carrace"))
	fmt.Println(isAnagram("jar", "jam"))

}
