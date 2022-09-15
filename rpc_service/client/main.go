package main

import (
	"fmt"

	proto "aa/proto"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {

	service := micro.NewService()

	service.Init()

	greeter := proto.NewGreeterClient("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "Mike",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
