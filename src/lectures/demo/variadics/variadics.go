package main

import "fmt"

func sum(nums ...int) int {
	// nums will be slice/array
	sum := 0

	for _, value := range nums {
		sum += value
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sum(all...)

	fmt.Println(answer)
}
