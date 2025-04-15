package pkg4for

import (
	"fmt"
)

func Func4For() {
	fmt.Println("Examples for FOR")

	fmt.Println("basic loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("break statement")
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("continue statement")
	b := 3
	for i := 0; i < 10; i++ {
		if i == b {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("while statement")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("repeat statement")
	i = 0 // reassigning
	for {
		fmt.Println(i)
		i = i + 1
		if i > 9 {
			break
		}
	}

	fmt.Println("range loop")
	s := "Hello, world!"
	for k, v := range s { // v is the rune
		fmt.Println(k, v, string(v))
	}

	fmt.Println("range loop for unicode")
	s = "👋 🌎"
	for k, v := range s { // v is the rune
		fmt.Println(k, v, string(v))
	}
}
