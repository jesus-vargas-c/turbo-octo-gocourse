package main

import (
	"fmt"
)

var a int

type dinero int

var b dinero

// main is the entry point for the program
func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
