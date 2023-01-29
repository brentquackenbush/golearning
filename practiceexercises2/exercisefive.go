package practiceexercises2

import "fmt"

/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
*/

func main() {
	a := `what
in
the 
world
is
going
on
in
here
"MEOW"`
	fmt.Println(a)
}
