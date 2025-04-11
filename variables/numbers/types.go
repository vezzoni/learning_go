package numbers

import "fmt"

func Example2() {
	fmt.Println("Examples 2")

	/*
		byte:  uint8
		 int:  int32 or  int64
		uint: uint32 or uint64
	*/

	var i int8 = 20
	var j int32 = 40

	fmt.Println(int32(i) + j) // type conversion
	fmt.Println(i + int8(j))

	var f float32 = 5.6
	fmt.Println(float32(i) + f)
	fmt.Println(i + int8(f))
	fmt.Println(i + int8(f+1.9))
}
