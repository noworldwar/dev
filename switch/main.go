package main

import "fmt"

var (
	i interface{}
)

func assert(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("%v is int(%T)\r\n", t, t)
	case int64:
		fmt.Printf("%v is int64(%T)\r\n", t, t)
	case int32:
		fmt.Printf("%v is int32(%T)\r\n", t, t)
	case string:
		fmt.Printf("%v is string(%T)\r\n", t, t)
	case float64:
		fmt.Printf("%v is float(%T)\r\n", t, t)
	default:
		fmt.Printf("Unknown type\r\n")
	}
}

func main() {
	assert(9000000000000000000)
	assert(0.1)
	assert("sgerger")
}
