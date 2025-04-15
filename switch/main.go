package main

import (
	"fmt"
	"os"
)

func main() {
	word := "hello"
	if len(os.Args) > 1 {
		word = os.Args[1]
	}

	// if & else chaining
	if word == "hello" {
		fmt.Println("Hi yourself")
	} else if word == "goodbye" {
		fmt.Println("So long!")
	} else if word == "greetings" {
		fmt.Println("Salutations!")
	} else {
		fmt.Println("I don't know what you said")
	}

	// switch
	switch word {
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye":
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said")
	}

	greet := "greetings"
	switch l := len(word); word {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough // <- only do the next case
	case "hello":
		fmt.Println("Hi yourself")
	case "farewell": // <- do nothing
	case "goodbye", "bye":
		fmt.Println("So long!")
	case greet:
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long")
	}

	c := "crackerjack"
	switch l := len(word); {
	case word == "hi":
		fmt.Println("Very informal!")
		fallthrough // <- only do the next case
	case word == "hello":
		fmt.Println("Hi yourself")
	case l == 1:
		fmt.Println("I don't know any one letter words")
	case 1 < l && l < 10, word == c: // <- exclusive OR
		fmt.Println("This word is either", c, "or it is 2-9 characters long")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long")
	}
}
