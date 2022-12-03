package main

import "fmt"

// func iterage(slice []int) {
// 	for i:=0;i<len(slice);i++{
// 		//..
// 	}
// }

// Slices are companion type that work with arrays
// They enable view into an array

// Functions can accept slices

// Basically we are *slicing* an array
// array == 1,2,3,4,5
// slice == 2,3,4

// we cannot access elements before starting position
// i.e. cannot access 1 here

// but we can access after the last position
// Slices points to array hence does not copy
// minimum memory
func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	// mySlice := []int{1, 2, 3}

	// item1 := mySlice[0]

	// slice[startIndex:endIndex] ---> [startIndex,endIndex)

	// numbers := [...]int{1, 2, 3, 4}

	// slice1 := numbers[:] --> same as array
	// slice2 := numbers[1:] --> start at 1 and everything upto the end
	// slice3 := slice2[:1] --> [2]

	// create arrays that can be extended

	// numbers := []int{1, 2, 3}
	// append(slice, ...numbers)
	// numbers = append(numbers, 4, 5, 6)

	// part1 := []int{1, 2, 3}
	// part2 := []int{5, 6, 4}

	// 3 dots can be used to extend a slice with another slice
	// part2... is spreading the slice
	// combined := append(part1, part2...)

	// Preallocation
	// make() is used to preallocate a slice
	// useful when number of elements is known,
	// but their values are still unknown

	// slice := make([]int, 10)

	// Multidimentional

	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	{"_", "_", "_"},
	// 	{"_", "_", "_"},
	// }

	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[0][0] = "X"
	// board[0][0] = "X"
	// board[0][0] = "X"

	route := []string{"Grocery", "Department store", "Salon"}
	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

}
