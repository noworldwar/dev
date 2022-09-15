package main

import (
	"fmt"
)

func main() {

	var mixed interface{}

	mixed = 123
	fmt.Println(mixed)

	mixed = "foo bar"
	fmt.Println(mixed)

	mixed = []string{"A", "B"}
	fmt.Println(mixed)

}
