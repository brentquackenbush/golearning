package main

//● Create a value and assign it to a variable.
//● Print the address of that value.

import (
	"fmt"
)

func main() {

	x := 2
	fmt.Println(&x)
}
