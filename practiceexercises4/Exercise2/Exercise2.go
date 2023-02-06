package main

/*
Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign 10 VALUES
● Range over the slice and print the values out.
● Using format printing
○ print out the TYPE of the slice
*/

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
