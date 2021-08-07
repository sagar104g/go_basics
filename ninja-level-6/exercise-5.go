// Hands-on exercise #5
// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
// code: https://play.golang.org/p/NGGikHNvHv
// video: 106

package main

import "fmt"

type square struct {
	side int
}
type circle struct {
	radius int
}

func (c circle) area() {
	fmt.Println(float64(c.radius) * float64(c.radius) * 3.14)
}
func (s square) area() {
	fmt.Println(s.side * s.side)

}
func main() {
	a := circle{
		radius: 11,
	}
	b := square{
		side: 11,
	}
	a.area()
	b.area()

}
