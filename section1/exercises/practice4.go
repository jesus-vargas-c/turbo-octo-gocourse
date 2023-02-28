package exercises

import (
	"fmt"
)

type myType int

var x4 myType

func Practice4() {
	fmt.Println(x4)
	fmt.Printf("El tipo de x es: %T\n", x4)
	x4 = 42
	fmt.Println(x4)

}
