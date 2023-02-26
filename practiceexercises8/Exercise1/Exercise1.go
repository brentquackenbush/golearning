package main

/**
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this
hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
*/
import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	//put new code here
	fmt.Println(users)

	bs, error := json.Marshal(users)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(bs))
}
