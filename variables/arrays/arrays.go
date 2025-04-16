package arrays

import "fmt"

func Example6() {
	var vals [3]int

	vals[0] = 2
	vals[1] = 4
	vals[2] = 6

	//var vals2 [4]int = vals <- in go, the size of an array makes part of its type

	fmt.Println(vals, vals[0], vals[1], vals[2])
}
