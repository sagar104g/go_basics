// Hands-on exercise #3
// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
// solution: https://play.golang.org/p/SGfiULXzAB
// video: 073

package main

import "fmt"

func main() {
	a := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range a[:5] {
		fmt.Printf("%v %t\n", v, v)
	}
	fmt.Printf("\n")

	for _, v := range a[5:] {
		fmt.Printf("%v %t\n", v, v)
	}
	fmt.Printf("\n")

	for _, v := range a[2:7] {
		fmt.Printf("%v %t\n", v, v)
	}
	fmt.Printf("\n")

	for _, v := range a[1:6] {
		fmt.Printf("%v %t\n", v, v)
	}
}
