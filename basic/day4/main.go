package main

import (
	"fmt"
)

type Person struct {
	name   string
	isPapa bool
}

type Member struct {
	name     string
	isLeader bool
}

type Human struct {
	name string
	age  int
}

type Item struct {
	item  string
	price float64
}

var findOldestPeson = func(s []Human) string {
	max_age := s[0].age
	var name string = s[0].name

	for i := 1; i < len(s); i++ {
		if s[i].age > max_age {
			max_age = s[i].age
			name = s[i].name
		}

	}
	return name

}

func findExpensiveItem(s []Item) string {
	max_price := s[0].price
	item_name := s[0].item

	for i := 1; i < len(s); i++ {
		if s[i].price > max_price {
			max_price = s[i].price
			item_name = s[i].item
		}
	}
	return item_name
}

func findLeader(s []Member) string {
	for _, value := range s {
		if value.isLeader {
			return value.name
		}
	}
	return "Not found"
}

func findParent(p []Person) string {
	for _, value := range p {
		if value.isPapa == true {
			return value.name
		}
	}
	return "Not found"
}

func main() {

	people := []Person{{"vivek", false}, {"ankur", false}, {"alex", true}}

	fmt.Println(findParent(people))

	team := []Member{{"vivek", false}, {"ankur", true}, {"alex", true}}

	fmt.Println(findLeader(team))

	humans := []Human{{"vivek", 73}, {"ankur", 85}, {"alex", 65}}

	fmt.Println(findOldestPeson(humans))

	items := []Item{{item: "Laptop", price: 999.99}, {item: "Phone", price: 699.99}, {item: "Tablet", price: 399.99}}
	fmt.Println(findExpensiveItem(items))

	items1 := []Item{{item: "Phone", price: 699.99}, {item: "Tablet", price: 399.99}}
	fmt.Println(findExpensiveItem(items1))

}
