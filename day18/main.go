package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}
type Addr struct {
	city string
	age  int
}
type Profile struct {
	Person
	Addr
}

func mergeObj(o1 Person, o2 Addr) Profile {
	p := Profile{Person: o1, Addr: o2}
	return p
}

func mergeMap(map1 map[string]interface{}, map2 map[string]interface{}) map[string]interface{} {
	map3 := make(map[string]interface{})

	for i, v := range map1 {
		map3[i] = v
	}
	for i, v := range map2 {
		map3[i] = v
	}
	return map3
}

func main() {

	p := Person{"vivek", 23}
	j := Addr{"bbk", 24}

	m1 := make(map[string]interface{})
	m1["a"] = 1
	m1["b"] = 2

	m2 := make(map[string]interface{})
	m2["b"] = 3
	m2["c"] = 4

	m4 := map[string]interface{}{"id": 101, "name": "Laptop", "price": 50000}
	m5 := map[string]interface{}{"fontSize": 14, "language": "en"}

	fmt.Println(mergeObj(p, j))
	fmt.Println(mergeMap(m1, m2))
	fmt.Println(mergeMap(m4, m5))
}
