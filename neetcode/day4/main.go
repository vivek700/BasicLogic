package main

import (
	"fmt"
)

type Key struct {
	Value [26]int
}

func isGroupAnagram(str []string) [][]string {

	hashmap := make(map[Key][]string)

	for _, v := range str {
		count := [26]int{}

		for _, s := range v {
			count[int(s)-int('a')] += 1
		}
		key := Key{Value: count}

		hashmap[key] = append(hashmap[key], v)
	}
	values := [][]string{}

	for _, v := range hashmap {
		values = append(values, v)
	}

	return values

}

func main() {
	str := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	str1 := []string{"x"}

	fmt.Println(isGroupAnagram(str))
	fmt.Println(isGroupAnagram(str1))
}
