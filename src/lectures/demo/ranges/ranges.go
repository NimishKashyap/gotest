package main

import "fmt"

func main() {
	// range creates an iterator

	slice := []string{"hello", "world", "!"}

	// Slice/array iterator index, element
	// individual letters

	// normal for loop actually iterates through each byte
	// so use range keyword while iterating strings
	for i, element := range slice {
		fmt.Println(i, element, ":")

		// glyph representation
		for _, ch := range element {
			fmt.Printf("    %q\n", ch)
		}
	}

}
