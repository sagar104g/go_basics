// Hands-on exercise #4
// Follow these steps:
// start with this slice
// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// append to that slice this value
// 52
// print out the slice
// in ONE STATEMENT append to that slice these values
// 53
// 54
// 55
// print out the slice
// append to the slice this slice
// y := []int{56, 57, 58, 59, 60}
// print out the slice
// solution: https://play.golang.org/p/QUYhtSBaDS
// video: 074

package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	for _, v := range x {
		fmt.Printf("%v %t\n", v, v)
	}
	x = append(x, 53, 54, 55)

	for _, v := range x {
		fmt.Printf("%v %t\n", v, v)
	}
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	for _, v := range x {
		fmt.Printf("%v %t\n", v, v)
	}
}