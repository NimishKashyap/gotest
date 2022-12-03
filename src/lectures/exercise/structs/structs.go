//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	s1, s2 int
}

func area(r *Rectangle) int {
	return r.s1 * r.s2
}

func perimeter(r *Rectangle) int {
	return 2 * (r.s1 + r.s2)
}
func main() {
	var r Rectangle = Rectangle{1, 2}
	fmt.Println(area(&r))
	fmt.Println(perimeter(&r))

	r.s1 = r.s1 * 2
	r.s2 = r.s2 * 2

	fmt.Println(area(&r))
	fmt.Println(perimeter(&r))

}
