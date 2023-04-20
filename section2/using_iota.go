package section2

import (
	"fmt"
)

const (
	a = 2023 + iota
	b = 2023 + iota
	c = 2023 + iota
	d = 2023 + iota
)

func iota_func() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
