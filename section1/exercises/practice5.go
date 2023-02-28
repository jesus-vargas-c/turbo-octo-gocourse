package exercises

import (
	"fmt"
)

type numero int

var x5 numero
var y5 int

func Practice5() {
	fmt.Println(x5)
	fmt.Printf("El tipo de x es: %T\n", x5)

	x5 = 42
	fmt.Println(x5)

	y5 = int(x5)
	fmt.Println(y5)
	fmt.Printf("%T", y5)
}
