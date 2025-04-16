package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	//rand.Seed(time.Now().UnixNano()) <- Seed is deprecated
	a := rand.Intn((10))
	b := rand.Intn((10))
	c := int(math.Max(float64(a), float64(b)))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, "is bigger")
}
