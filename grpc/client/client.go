package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "server/proto"
)

func main() {
	// 連線到遠端 gRPC伺服器
	conn, err := grpc.Dial("loaclhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線錯誤: %v", err)
	}
	defer conn.Close()

	// 建立新的 Calculator客戶端，用來使用 Calculator的所有方法
	c := pb.NewCalculatorClient(conn)

	// 傳送新請求到遠端 gRPC 伺服器Calculator中，呼叫 Plus
	r, err := c.Plus(context.Background(), &pb.CalcRequest{NumA: 16, NumB: 28})
	if err != nil {
		log.Fatalf("無法執行 Plus: %v", err)
	}
	log.Printf("回傳結果 : %d", r.Result)
}