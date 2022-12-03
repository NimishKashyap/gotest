package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

// main package launches the program
func main() {

	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("An exciting message")
}
