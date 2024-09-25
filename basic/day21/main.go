package main

import "fmt"

func calculateCost(s []Cart) int {
	total := 0
	for _, v := range s {
		total += v.Cost()

	}
	return total
}

type Cart struct {
	price    int
	quantity int
}

func (item Cart) Cost() int {
	return item.price * item.quantity

}

func main() {

	items := []Cart{{price: 100, quantity: 2}, {price: 200, quantity: 1}, {price: 50, quantity: 4}}
	items1 := []Cart{{price: 30, quantity: 3}, {price: 15, quantity: 2}}
	items2 := []Cart{{price: 200, quantity: 3}, {price: 50, quantity: 5}}
	items3 := []Cart{{price: 10, quantity: 7}, {price: 5, quantity: 4}}
	groceryList := []Cart{{price: 20, quantity: 2}, {price: 15, quantity: 3}}
	groceryList1 := []Cart{{price: 7, quantity: 5}, {price: 3, quantity: 8}}
	invoice := []Cart{{price: 50, quantity: 10}, {price: 25, quantity: 4}}
	invoice1 := []Cart{{price: 100, quantity: 1}, {price: 40, quantity: 2}}

	fmt.Println(calculateCost(items))
	fmt.Println(calculateCost(items1))
	fmt.Println(calculateCost(items2))
	fmt.Println(calculateCost(items3))
	fmt.Println(calculateCost(groceryList))
	fmt.Println(calculateCost(groceryList1))
	fmt.Println(calculateCost(invoice))
	fmt.Println(calculateCost(invoice1))

}
