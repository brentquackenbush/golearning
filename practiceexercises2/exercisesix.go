package practiceexercises2

import (
	"fmt"
)

/*
Hands-on exercise #6
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

const (
	v = 2023 + iota
	x = 2023 + iota
	y = 2023 + iota
	z = 2023 + iota
)

func main() {
	fmt.Println(v)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
