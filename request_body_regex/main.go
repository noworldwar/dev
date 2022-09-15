package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	regexPattern := "Content(.*?)\\n"
	b, err := ioutil.ReadAll(file)
	re := regexp.MustCompile(regexPattern)
	result := re.FindAllString(string(b), -1)
	fmt.Println(result)
}
