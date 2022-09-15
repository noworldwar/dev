package main

import (
	"fmt"
)

func main() {

	result := multiply("5678", "1234")
	fmt.Println(result)

}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := make([]byte, len(num1)+len(num2))
	for i := range ans {
		ans[i] = '0'
	}
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := (num1[i] - '0') * (num2[j] - '0')
			tmp := ans[i+j+1] - '0' + product
			ans[i+j+1] = tmp%10 + '0'
			ans[i+j] = (ans[i+j] - '0' + tmp/10) + '0'
		}
	}
	if ans[0] == '0' {
		return string(ans[1:])
	}
	return string(ans)
}
