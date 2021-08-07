// Hands-on exercise #10
// Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
// solution: https://play.golang.org/p/TYl5EbjoeC
// video: 080

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
	value, status := m["aabb"]
	fmt.Println(value, status)

	if value, status = m["bond_james"]; status {
		fmt.Println("before delete ", value, status)
	}
	delete(m, "bond_james")
	fmt.Println("after delete ", m["bond_james"])
}
