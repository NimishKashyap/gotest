package main

import "fmt"

// -er as suffix
type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}

func (c Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("Add dressing")
}

func prepareDish(dishes []Preparer) {
	fmt.Println("Prepare dishes")

	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Winds"), Salad("Iceberg Salad")}

	prepareDish(dishes)
}
