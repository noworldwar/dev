package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	questions := [3]string{
		"Please enter betlimit (TEST-A,TEST-B,TEST-C)",
		"Please enter iv value (The first 16 words of appSecret)",
		"Please enter Secret Key (Secret Key after encrypted appSecret by MD5)",
	}

	var vals []string
	var val string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		vals = []string{}
		fmt.Println(questions[0])
		for i := 1; scanner.Scan(); i++ {
			val = scanner.Text()
			if strings.TrimSpace(val) == "" {
				i--
			} else {
				// fmt.Println(val)
				vals = append(vals, strings.TrimSpace(val))
			}

			if i == 3 {
				break
			}
			fmt.Println(questions[i])
		}
		result := Aes256Encode(vals[0], vals[2], vals[1], aes.BlockSize)
		fmt.Println()
		fmt.Println()
		fmt.Println("Output: ", result)
		fmt.Println("Encode to URL-encoded format: ", url.QueryEscape(result))
		fmt.Println()
		time.Sleep(2 * time.Second)
	}

}

func Aes256Encode(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return ""
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
