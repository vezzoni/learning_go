package main

import (
	"fmt"

	leftpad "github.com/vezzoni/learning_go/packages"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '🙂'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '🙂'))
}
