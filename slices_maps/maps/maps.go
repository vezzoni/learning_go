package maps

import "fmt"

func modMap(m map[string]int) {
	m["cheese"] = 20
}

func Example4Maps() {
	fmt.Println("Example for Maps")
	// java: hashmap
	// python: dict

	m := make(map[string]int) // creates a map: key string & value int
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("hello in m:", h)
	fmt.Println("a in m:", m["a"]) // there's no exist

	if v, ok := m["hello"]; ok {
		fmt.Println("hello in m:", v)
	}

	m["hello"] = 20
	fmt.Println("hello in m:", m["hello"])
	fmt.Println("hello in h:", h) // didn't change

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50,
	} // another syntax to create a map

	for k, v := range m2 { // the order of the iteration is random
		fmt.Printf("%s key has %d value\n", k, v)
	}

	fmt.Println("b in m2:", m2["b"])
	delete(m2, "b")
	fmt.Println("b in m2:", m2["b"])

	m = map[string]int{
		"a": 1,
		"b": 2,
	}
	var m3 map[string]int // just declaration: nil
	fmt.Println("m3:", m3)
	//m3["foo"] = 17 <- panic

	fmt.Println("goodbye in m:", m["goodbye"])
	m3 = m // they share the same location at the memory
	m3["goodbye"] = 400
	fmt.Println("goodbye in m3:", m3["goodbye"])
	fmt.Println("goodbye in m:", m["goodbye"])

	modMap(m)
	fmt.Println("cheese in m:", m["cheese"])
}
