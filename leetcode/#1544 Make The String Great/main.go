package main

import "fmt"

func main() {
	s := "leEeetcode"

	flag := false
	for {
		if flag == true || s == "" {
			break
		}
		s, flag = makeGood(s)
	}
	fmt.Println(s)
}

func makeGood(s string) (string, bool) {
	if len(s) <= 1 {
		return s, true
	}
	flag := false
	for i := 1; i < len(s); i++ {
		if int(s[i])-int(s[i-1]) == 32 || int(s[i-1])-int(s[i]) == 32 {
			return s[:i-1] + s[i+1:], false
		}
		flag = true
	}
	return s, flag
}
