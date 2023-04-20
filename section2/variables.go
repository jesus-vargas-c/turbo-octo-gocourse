package section2

import (
	"fmt"
)

var decimal_value int = 42

func PrintValue() {
	fmt.Println("values printing")
	fmt.Printf("%d\t%b\t%x", decimal_value, decimal_value, decimal_value)
}
