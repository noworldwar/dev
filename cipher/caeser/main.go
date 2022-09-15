package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	output := Encrypt(text, 5)
	fmt.Println("Encrypt Text : ", output)
	fmt.Println("Decrypt Text : ", Decrypt(output, 5))
}

func Encrypt(input string, key int) string {

	key8 := byte(key%26+26) % 26

	var outputBuffer []byte

	for _, r := range input {
		newByte := byte(r)
		if 'A' <= r && r <= 'Z' {
			outputBuffer = append(outputBuffer, 'A'+(newByte-'A'+key8)%26)
		} else if 'a' <= r && r <= 'z' {
			outputBuffer = append(outputBuffer, 'a'+(newByte-'a'+key8)%26)
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}
	}

	return string(outputBuffer)
}

func Decrypt(input string, key int) string {
	return Encrypt(input, 26-key)
}
