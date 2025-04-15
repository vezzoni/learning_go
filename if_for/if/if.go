package pkg4if

import (
	"fmt"
)

func Func4If() {
	fmt.Println("Examples for IF")
	a := 10

	if a > 5 {
		b := a / 2 // <- only exist into this scope
		fmt.Println("a is bigger than 5:", a, b)
	} else {
		fmt.Println("a is less than or equal to 5")
	}

	//fmt.Println(b) // <- out of scope

	if b := a / 2; b > 5 { // brand new declare of b
		fmt.Println("b is bigger than 5:", a, b)
	} else {
		fmt.Println("b is less than or equal to 5", a, b)
	}

	//fmt.Println(b) // <- out of scope
}
