package main

import "fmt"

/*
For this exercise
1. Create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER “o” using the “VAR”
keyword
3. in func main
a. print out the value of the variable “o”
b. print out the type of the variable “o”
c. assign 42 to the VARIABLE “o” using the “=” OPERATOR
d. print out the value of the variable “o”
*/

type meow int

var o meow

func main() {
	fmt.Println(o)
	fmt.Printf("%T\n", o)
	o = 42
	fmt.Println(o)
}
