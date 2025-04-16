package slices

import "fmt"

func modSlice(s []string) {
	s[0] = "hello"
}

func Example4Slices() {
	fmt.Println("Example for Slices")

	s := make([]string, 0) // creating a slice using a built-in make function. used to create instances. slice length 0
	fmt.Println("length of s:", len(s))
	s = append(s, "hello") // increasing the size of the slice, adding the new element at the end
	fmt.Println("length of s:", len(s))
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0])

	s2 := make([]string, 2)
	fmt.Println("contents of s2[0]:", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))

	s3 := []string{"a", "b", "c"} // creates a slice with specified elements

	for i, v := range s3 {
		fmt.Printf("%d index has the %s value\n", i, v)
	}

	s4 := s3[0:2] // slice of slice: they share the same location at the memory.
	fmt.Println("s4:", s4)
	s3[0] = "d" // s4 points to the same location at the memory
	fmt.Println("s4:", s4)

	var s5 []string // <- nil: no reference into the memory
	s5 = s3         // they share the same location at the memory
	s5[1] = "camel"
	fmt.Println("s3:", s3)

	modSlice(s3)
	fmt.Println("s3[0]:", s3[0])

	uniHello := "👋 🌍"
	bytes := []byte(uniHello) // slice of bytes
	fmt.Println("from bytes:", bytes)
	runes := []rune(uniHello)
	fmt.Println("from runes:", runes)
	runes[1] = 'a' // code 97
	fmt.Println("from runes:", runes)
	fmt.Println("from hello:", uniHello)
}
