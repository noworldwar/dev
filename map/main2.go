package main

import "fmt"

type betData struct {
	betTime string
}

func main() {
	userData := make(map[int]*betData)

	// userData[0] = betData{
	// 	betTime: "123",
	// }
	// userData := map[int]*betData{}
	userData[0] = &betData{}
	userData[0].betTime = "123"
	fmt.Println(userData[0])
}
