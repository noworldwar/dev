package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {

	fmt.Println(Aes256Decode("pLw9hkxtzS+ldgZBfTCzXA==", "c84b812db9f1900e2abb01d770973720", "iRgUdd5J55124pbV"))

}

func Aes256Decode(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Printf("DecodeString %v | %v err: %v", cipherText, cipherTextDecoded, err)
		return ""
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		fmt.Printf("NewCipher %v | %v err: %v", cipherText, cipherTextDecoded, err)
		return ""
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return strings.TrimSpace(string(cipherTextDecoded))
}
