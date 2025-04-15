package main

import "fmt"

func setTo10(a *int) {
	*a = 10
}

func setTo10Fail(a *int) {
	a = new(int) // creates a new pointer. the "a" parameter is left alone
	*a = 10
}

func main() {

	a := 10
	b := &a // assigning the reference (memory address where the "a" variable value is stored)
	c := a

	fmt.Println(a, b, *b, c)

	a = 20 // the variable b points to the same memory address, so the variable b has the same value
	fmt.Println(a, b, *b, c)

	*b = 30 // assigning a new value to the target reference (using the memory address of the variable b, which is the same of the variable a)
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	var i *int // zero value for a pointer -> nil (absence of a pointer value)
	fmt.Println(i, "from i pointer")
	//fmt.Println(*i, "from i pointer") // nil. not initialized (there's no memory address): panic

	t := new(int) // makes a pointer for the type -> !nil
	fmt.Println(t, "from t pointer")
	fmt.Println(*t, "from t pointer")

	x := 20
	fmt.Println(x, "from x pointer")
	setTo10(&x)
	fmt.Println(x, "from x pointer")
	x = 20
	fmt.Println(x, "from x pointer")
	setTo10Fail(&x)
	fmt.Println(x, "from x pointer")
}
