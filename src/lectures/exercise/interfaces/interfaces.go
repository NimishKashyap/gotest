//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lifter interface {
	Lift(int) int
}

type Motorcycle struct {
	size int
}
type Cars struct {
	size int
}
type Trucks struct {
	size int
}

func (m Motorcycle) Lift(size int) int {
	return size + 2
}
func (m Cars) Lift(size int) int {
	return size + 5
}
func (m Trucks) Lift(size int) int {
	return size + 9
}

func Lift(l []Lifter) {
	for _, value := range l {
		var val int
		if vehicle, ok := value.(Motorcycle); ok {
			val = vehicle.Lift(5)
		} else if vehicle, ok := value.(Cars); ok {
			val = vehicle.Lift(9)
		} else {
			val = vehicle.Lift(15)
		}

		fmt.Println(val)

	}
}

func main() {
	vehicles := []Lifter{Motorcycle{size: 10}, Cars{size: 15}, Trucks{size: 30}}

	Lift(vehicles)
}
