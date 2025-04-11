package numbers

import "fmt"

func Example1() {
	fmt.Println("Examples 1")

	var i int = 10 // declared variables must be used otherwise there's a compile error
	fmt.Println(i)

	var k = 20 // type inference
	fmt.Println(k)

	t := 30 // shorted-hand syntax can't be used outside from a function scope
	fmt.Println(t)

	var b byte = 40
	fmt.Println(b)

	var x int
	fmt.Println(x) // prints 0

	x = 50 // reassignment
	fmt.Println(x)
}
