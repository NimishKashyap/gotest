package main

import "fmt"

type Counter struct {
	hits int
}

// with structers, no need to use dereference *
func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World"

	fmt.Println(hello, world)
	replace(&hello, "hi", &counter)

	fmt.Println(hello, world)

	phrase := []string{hello, world}

	fmt.Println(phrase)

	replace(&phrase[1], "GO!", &counter)

	fmt.Println(phrase)

	fmt.Println(hello, world)

}
