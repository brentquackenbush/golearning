package main

/*
Create a for loop using this syntax
● for condition { }
Have it print out the years you have been alive.
*/

import (
	"fmt"
)

func main() {
	birthday := 1997
	for birthday < 2023 {
		fmt.Println(birthday)
		birthday++
	}
}
