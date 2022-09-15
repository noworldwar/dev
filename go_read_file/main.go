package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	data, err := ioutil.ReadFile("router.go")
	if err != nil {
		fmt.Println("File Reading Error", err)
		return
	}
	fmt.Println(string(data))
}
