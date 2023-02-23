package main

import "fmt"

type Counter struct {
	hits int
}

type Operation interface {
	increment()
	decrement()
}

// with structers, no need to use dereference *
func (counter *Counter) increment() {
	counter.hits += 1
}

func (counter *Counter) decrement() {
	counter.hits -= 1
}

// func replace(old *string, new string, counter *Counter) {
// 	*old = new
// 	increment(counter)
// }

func display(obj Operation) {
	fmt.Println(obj)
}

func main() {
	counter := Counter{}

	c2 := 1
	// decrement(&c2)
	fmt.Println(c2)

	// hello := "Hello"
	// world := "World"

	// fmt.Println(hello, world)
	// replace(&hello, "hi", &counter)

	// fmt.Println(hello, world)

	// phrase := []string{hello, world}

	// fmt.Println(phrase)

	// replace(&phrase[1], "GO!", &counter)

	// fmt.Println(phrase)

	// fmt.Println(hello, world)
	display(counter)

}
