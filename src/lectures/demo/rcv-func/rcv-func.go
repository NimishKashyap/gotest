package main

import "fmt"

// type Coord struct {
// 	X, Y int
// }

// func (coord *Coord) shift(x, y int) {
// 	coord.X += x
// 	coord.Y += y
// }

// func (c Coord) shiftDist(other Coord) Coord {
// 	return Coord{other.X - c.X, other.Y - c.Y}
// }

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	// coord := Coord{5, 5}
	// second := Coord{1, 5}

	lot := ParkingLot{spaces: make([]Space, 10)}

	fmt.Println(lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupation:", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)

	// // coord.shift(1, 1)

	// distance := coord.shiftDist(second)

	// fmt.Println(distance)
}
