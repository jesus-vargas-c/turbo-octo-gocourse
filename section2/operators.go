package section2

import (
	"fmt"
)

func Comparisons() {
	casea := 10 == (1 + 9)
	caseb := 10 <= (1 + 9)
	casec := 11 >= (10 + 1)
	cased := "a" != "b"
	casee := 1 < 2
	casef := 3 > 2

	fmt.Println(casea)
	fmt.Println(caseb)
	fmt.Println(casec)
	fmt.Println(cased)
	fmt.Println(casee)
	fmt.Println(casef)
}
