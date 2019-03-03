package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(999)
	fmt.Println("My favorite number is", rand.Intn(20))
}
