package main

import "fmt"

func findItem(s []string, keyword string) int {
	l := 0
	h := len(s) - 1

	for l <= h {
		m := (l + h) / 2
		if keyword == s[m] {
			return m
		} else if keyword > s[m] {
			l = m + 1

		} else {
			h = m - 1
		}

	}
	return -1

}

func main() {
	books := []string{"A Tale of Two Cities", "Great Expectations", "Moby Dick", "Pride and Prejudice"}
	books1 := []string{"A Brief History of Time", "Harry Potter", "The Hobbit", "War and Peace"}
	books2 := []string{"1984", "Brave New World", "To Kill a Mockingbird", "Wuthering Heights"}

	fmt.Println(findItem(books, "Moby Dick"))
	fmt.Println(findItem(books, "Great Expectations"))
	fmt.Println(findItem(books, "Great"))
	fmt.Println(findItem(books1, "Harry Potter"))
	fmt.Println(findItem(books2, "To Kill a Mockingbird"))
	fmt.Println(findItem(books2, "Harry Potter"))
	fmt.Println(findItem(books2, "Brave New World"))
}
