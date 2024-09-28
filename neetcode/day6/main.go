package main

import (
	"fmt"
	"log"
	"strconv"
)

func encode(strs []string) string {

	newStr := ""
	for _, v := range strs {
		s := strconv.Itoa(len(v))
		ns := "[" + s + "]"
		newStr = newStr + ns + v
	}

	return newStr

}

func decode(str string) []string {
	newstr := ""
	newStrSlice := []string{}

	IndexOfLen := 0
	length := 0
	for i := 0; i < len(str); i++ {

		if i == IndexOfLen {
			nstr := ""
			if string(str[i]) == "[" {
				i++
				IndexOfLen++
				for string(str[i]) != "]" {
					nstr += string(str[i])
					i++
					IndexOfLen++
				}
			}

			value, err := strconv.Atoi(string(nstr))
			if err == nil {
				length = value
				continue

			} else {
				log.Fatal("err:", err)
			}
		}

		newstr = fmt.Sprintf("%s%s", newstr, string(str[i]))

		if length == len(newstr) {

			newStrSlice = append(newStrSlice, newstr)
			newstr = ""
			IndexOfLen += length + 1
			length = 0

		}

	}

	return newStrSlice
}
func main() {

	input := []string{"neet", "code", "love", "you"}
	input1 := []string{"we", "say", ":", "yes"}
	input2 := []string{"we", "say", ":", "yes", "!@#$%^&*()"}
	input3 := []string{"we", "say", ":", "yes", "!@#$%^&*()", "asfaklsfjlkasjgksjg", "aksfjkasfakldhfjkdsjlkgfasjgslkjkdsfjkasjfkasjf"}

	fmt.Println(encode(input))
	fmt.Println(encode(input1))
	fmt.Println(encode(input2))
	fmt.Println(encode(input3))
	fmt.Println(decode(encode(input)))
	fmt.Println(decode(encode(input1)))
	fmt.Println(decode(encode(input2)))
	fmt.Println(decode(encode(input3)))

}
