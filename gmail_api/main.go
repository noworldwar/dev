package main

import (

	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){

	url := "https://www.googleapis.com/gmail/v1/users/me/messages/"
	access_token := "?access_token=ya29.a0Adw1xeXvm2xOO7OehRjRt6_sxeCYGYOg86wjB3UdXv3wiwQWGNQsGRs_m83RY4LhED2lUGNgNiDH-nAA9OhbhW357_2fMVwUJy1t5HqYIiCI0AXSFkXrI009BzhTjQR-SyWNuYXwZq5WMMXtf7SpN_WxWOUstUW9yhI"
	resp, err := http.Get(url+access_token)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	s:= string(body)
	fmt.Println(s)
}