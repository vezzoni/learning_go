package main

import (
	"fmt"
	"language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
}
