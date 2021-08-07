// Hands-on exercise #9
// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
// solution: https://play.golang.org/p/_CkxAhRrDJ
// video: 079

package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james":      {`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range m {
		fmt.Printf("%v %v\n", k, v)
	}
}
