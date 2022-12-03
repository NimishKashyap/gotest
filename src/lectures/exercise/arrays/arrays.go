//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printLast(shoppingList [4]Product) {
	for i := 0; i < len(shoppingList); i++ {
		if i == 2 {
			fmt.Println(shoppingList[i])
		}
	}
}
func totalCost(shoppingList [4]Product) {
	cost := 0
	for i := 0; i < len(shoppingList); i++ {
		cost += shoppingList[i].price
	}
	fmt.Println(cost)
}

func totalItems(shoppingList [4]Product) {
	fmt.Println(len(shoppingList))

}

func main() {
	var shoppingList [4]Product = [4]Product{{name: "Apple", price: 10}, {name: "Mango", price: 20}, {name: "Guava", price: 15}}

	fmt.Println(shoppingList)

	printLast(shoppingList)
	totalCost(shoppingList)
	totalItems(shoppingList)

}
