//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import (
	"fmt"
)

type Items struct {
	item        string
	isActivated bool
}

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = true
}
func deactivate(tag *SecurityTag) {
	*tag = false
}
func toggleActivation(item *Items) {
	item.isActivated = !item.isActivated

func checkout(items []Items) {
	for index := range items {
		items[index].isActivated = false
	}

}

func checkout2(items []Item) {
	fmt.Println("Checking out...")
	
	for index := range items {
		deactivate(&items[i].tag)
	}

}

func main() {
	items := []Items{{item: "1", isActivated: true}, {item: "2", isActivated: true}, {item: "3", isActivated: true}, {item: "4", isActivated: true}}

	toggleActivation(&items[0])

	fmt.Println(items)

	checkout(items)

	fmt.Println(items)
}
