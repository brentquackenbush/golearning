package main

/*
Using the code from the previous example,
use SLICING to create the following new slices
which are then printed:
● [4243444546] ● [4748495051] ● [4445464748] ● [4344454647]
*/

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
