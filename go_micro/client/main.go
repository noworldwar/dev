package main

import (
	"context"
	"fmt"

	// 引用上面生成的proto文件
	proto "aaa/proto"

	"github.com/micro/go-micro"
)

func main() {
	// new a service
	service := micro.NewService()

	// 初始化
	service.Init()

	// 建一個客戶端
	cl := proto.NewGreeterService("srv.greeter", service.Client())

	// 送出請求
	rsp, err := cl.Hello(context.Background(), &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
