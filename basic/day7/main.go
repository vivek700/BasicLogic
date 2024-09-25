package main

import (
	"fmt"
	"strings"
)

func includes(s []string, keyword string) []string {
	var slice []string
	for _, v := range s {
		found := strings.Contains(v, keyword)
		if found {
			slice = append(slice, v)
		}
	}
	return slice
}

func searchTitle(s []string, keyword string) []string {
	var slice []string

	for _, v := range s {
		found := strings.Contains(v, keyword)
		if found {
			slice = append(slice, v)
		}
	}
	return slice
}

func findContact(s []string, contact string) []string {
	var slice []string

	for _, v := range s {
		found := strings.Contains(v, contact)
		if found {
			slice = append(slice, v)
		}
	}
	return slice
}

func findIngredient(s []string, ingredient string) string {

	for _, v := range s {
		found := strings.Contains(v, ingredient)
		if found {
			return "Found"
		}
	}
	return "Not found"
}

func main() {

	messages := []string{"Hey, how are you?", "I miss you!", "Let's catch up soon!", "Good morning!"}
	messages1 := []string{"Hey, how are you?", "I miss you!", "Good morning!"}

	movies := []string{"Dilwale Dulhania Le Jayenge", "3 Idiots", "Lagaan", "Sholay"}
	movies1 := []string{"Dilwale Dulhania Le Jayenge", "Lagaan"}

	fmt.Println(includes(messages, "catch"))
	fmt.Println(includes(messages1, "hello"))

	fmt.Println(searchTitle(movies, "3 Idiots"))
	fmt.Println(searchTitle(movies1, "Sholay"))

	recipe := []string{"Cheeni", "Dhahi", "Aata"}

	fmt.Println(findIngredient(recipe, "Cheeni"))
	fmt.Println(findIngredient(recipe, "jeera"))

	phonebook := []string{"Amit", "Neha", "Ravi", "Sita"}

	fmt.Println(findContact(phonebook, "Neha"))
	fmt.Println(findContact(phonebook, "vivek"))
}
