package main

import "fmt"

/**
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

type person struct {
	first        string
	last         string
	favoriteFlav []string
}

func main() {
	testData := person{
		first: "Brent",
		last:  "Quack",
		favoriteFlav: []string{
			"vanilla",
			"chocolate",
		},
	}

	for i, v := range testData.favoriteFlav {
		fmt.Println(i, v)
	}

}
