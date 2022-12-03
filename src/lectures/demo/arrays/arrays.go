package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	// var myarray [3]int
	// myarray := [3]int{3,2,6}
	// myarray := [...]int{3,5,7}

	// myarray := [4]int{3, 5, 7}

	// Accessing data
	// item := myarray[0]

	// myarray := [...]int{4, 3, 2}
	// var myarray [3]int = [3]int{2, 3, 5}

	// for i := 0; i < len(myarray); i++ {
	// 	// copy
	// 	item := myarray[i]
	// 	fmt.Println(item)
	// }

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}
	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)
}
