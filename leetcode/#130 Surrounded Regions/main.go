package main

import "fmt"

func main() {
	inputStr := [][]string{{"X", "X", "X", "X"}, {"X", "O", "O", "X"}, {"X", "X", "O", "X"}, {"X", "O", "X", "X"}}

	output := solve(inputStr)
	for _, row := range output {
		fmt.Println(row)
	}
}

func solve(board [][]byte) [][]byte {

}
