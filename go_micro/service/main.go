package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	// 引用上面生成的proto文件
	greeter "aaa/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	//new一個service
	service := micro.NewService(
		micro.Name("srv.greeter"),
		micro.Version("latest"),
	)

	//解析命令行
	service.Init()

	//註冊 handler
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	//啟動service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
