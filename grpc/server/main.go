package main

import (
	"log"
	"net"

	pb "server/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server 建構體會實作 Calculator 的 gRPC 伺服器
type server struct{}

// 加總函數
func (s *server) Plus(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {

	// 計算
	result := in.NumA + in.NumB

	// 包裝成 Protobuf 建構體並回傳
	return &pb.CalcReply{Result: result}, nil
}

func (s *server) Subtraction(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {

	// 計算
	result := in.NumA - in.NumB

	// 包裝
	return &pb.CalcReply{Result: result}, nil

}

func main() {

	// 監聽指定埠口
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("無法監聽該埠口 : %v", err)
	}

	// 建立新 gRPC 伺服器並註冊 Calulator 服務
	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &server{})

	// 在gRPC 伺服器上註冊反射服務
	reflection.Register(s)
	// 開始在指定埠口服務
	if err := s.Serve(lis); err != nil {
		log.Fatalf("無法提供服務: %v", err)
	}
}
