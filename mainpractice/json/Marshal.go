package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Brent",
		Last:  "Junior",
		Age:   23,
	}
	p2 := person{
		First: "Taco",
		Last:  "Spice",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	//to send data (encode) into json. It is called Marshalling in Go

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
