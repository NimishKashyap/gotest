package main

import "fmt"

func main() {
	// var a [5]int = [5]int{1, 2, 3, 4, 5}

	// var slice []int = a[0:4]

	// fmt.Println("slice", slice)
	// slice[0] = 100

	// fmt.Println(a)

	// var another []int
	// another = append(another, 1, 2, 3, 4, 5)
	// fmt.Println("another", another)

	// a := []int{1, 2, 3}

	// fmt.Print(a)

	// arr := [5]int{1, 2, 3, 4, 5}

	// sli := arr[:]

	// anothersli := []int{4, 5, 4}

	// fmt.Print(sli)

	// sli = append(sli, 5, 6, 7, 8)

	// sli = append(sli, anothersli...)

	// fmt.Println("slice", sli)
	// fmt.Println("arr", arr)

	demo := make([]int, 10)
	fmt.Println("demo", demo)
	fmt.Println("length", len(demo))

	board := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	fmt.Println(board)

}
