package main

import "fmt"

type Gender int

const (
	Male Gender = iota
	Female
)

type Person struct {
	name   string
	gender Gender
	age    uint
}

func main() {

	me := Person{gender: Male, age: 30, name: "Michael Chang"}
	fmt.Println(me)
}
