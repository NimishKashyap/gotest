package main

import "fmt"

// Key-value pair
// extremely high performance

func main() {

	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)
	// myMap := make(map[string]int)
	fmt.Println(len(shoppingList))

	delete(shoppingList, "milk")

	fmt.Println(shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]

	if !found {
		fmt.Println("Cereal is not found")
	} else {
		fmt.Println("Cereal boxes:", cereal)
	}

	totalItems := 0

	for _, value := range shoppingList {
		totalItems += value
	}

	fmt.Println("Total items required", totalItems)

	// myMap := map[string]int{
	// 	"item1": 1,
	// 	"item2": 2,
	// 	"item3": 3,
	// 	"item4": 4,
	// }
	// fmt.Println(myMap)

	// // insertion
	// myMap["favourite number"] = 5

	// // read
	// fav := myMap["favourite number"]
	// missing := myMap["age"] //default value is 0

	// // deleting
	// delete(myMap, "favourite number")

	// // check existence

	// fmt.Println(fav, missing)

	// // if !found {
	// // 	fmt.Println("Price not found")
	// // 	return
	// // }

	// // iteration

	// // key and value while using range with map
	// for key, value := range myMap {
	// 	fmt.Println("Key: ", key, "Value:", value)
	// }

}
