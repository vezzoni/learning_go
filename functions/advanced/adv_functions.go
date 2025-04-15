package advanced

import "fmt"

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperations(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makeAdder(b int) func(int) int {
	return func(a int) int { // <- closure: it keeps the state of b (the parameter)
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int { // <- closure: it keeps the state of f (the function)
		b := f(a)
		return b * 2
	}
}

func Example2() {
	fmt.Println("Example 2")

	myAddOne := addOne

	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1))

	// anonymous function
	myAddOne = func(a int) int {
		return a + 1
	}

	fmt.Println(myAddOne(1))

	printOperations(1, addOne)
	printOperations(1, addTwo)

	/*
		closure:
		local functions that access variables existent in the environment it was declared
	*/
	b := 2
	myAddOne = func(a int) int {
		return a + b
	}

	fmt.Println(myAddOne(1), "from myAddOne(1)")

	b = 2
	myNewAddOne := func(a int) {
		b = a + b
	}

	myNewAddOne(1)
	fmt.Println(b, "from myNewAddOne(1)")

	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fmt.Println(addOne(1), "from makeAdder(1)")
	fmt.Println(addTwo(1), "from makeAdder(2)")

	doubleAddOne := makeDoubler(addOne)
	fmt.Println(doubleAddOne(1), "from makeDoubler(addOne)")
}
