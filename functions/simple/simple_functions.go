package simple

import "fmt"

func addNumber(a int, b int) {
	fmt.Println(a + b)
}

// no overloading
// func addNumber(a int, b int, c int) {
// 	fmt.Println(a + b + c)
// }

func addNumbers(a int, b int) int {
	return a + b
}

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

// by value
func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s

	fmt.Println("in doubeFail:", a, arr, s)
}

func Example1() {
	fmt.Println("Example 1")

	addNumber(2, 3)
	addNumber(4, 10)
	addNumber(100, -100)

	a := addNumbers(2, 3)
	fmt.Println(a)
	a = addNumbers(4, 10)
	fmt.Println(a)
	a = addNumbers(100, -100)
	fmt.Println(a)

	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)
	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)
	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)
	divAndRemainder(-1, 20)

	a = 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFail(a, arr, s)
	fmt.Println("in main:", a, arr, s)
}
