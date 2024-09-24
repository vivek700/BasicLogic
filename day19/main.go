package main

import (
	"fmt"
	"strings"
)

func capitalize(v string) string {
	return strings.Replace(v, string(v[0]), strings.ToUpper(string(v[0])), 1)
}

func toUpperCase(s string) string {
	sp := strings.Split(s, " ")
	var upperS []string
	for _, v := range sp {
		upperS = append(upperS, capitalize(v))

	}
	return strings.Join(upperS, " ")
}

func toUpperCaseArrayItems(s []string) []string {
	var newSlice []string
	for _, v := range s {
		newSlice = append(newSlice, capitalize(v))
	}
	return newSlice

}

func main() {
	sentence := "openai is revolutionizing ai technology"
	sentence1 := "hello world"
	title := "the great gatsby"
	title1 := "harry potter and the philosopher's stone"
	headline := "breaking news: local hero saves the day"
	headline1 := "sports update: team india wins the match"

	names := []string{"rahul", "anjali", "vivek"}
	names1 := []string{"amit", "sneha", "piyush"}

	fmt.Println(toUpperCase(sentence))
	fmt.Println(toUpperCase(sentence1))
	fmt.Println(toUpperCase(title))
	fmt.Println(toUpperCase(title1))
	fmt.Println(toUpperCase(headline))
	fmt.Println(toUpperCase(headline1))

	fmt.Println(toUpperCaseArrayItems(names))
	fmt.Println(toUpperCaseArrayItems(names1))
}
