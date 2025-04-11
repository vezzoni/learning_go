package chars

import (
	"fmt"
)

func Example4() {
	fmt.Println("Examples 4	")

	var vals [3]int

	vals[0] = 2
	vals[1] = 4
	vals[2] = 6

	//var vals2 [4]int = vals // compile error because in go, the size of an array makes part of its type

	fmt.Println(vals, vals[0], vals[1], vals[2])
}
