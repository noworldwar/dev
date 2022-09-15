package main

import (
	"context"
	"fmt"

	. "github.com/PGITAb/bc-proto-entity-playerloginlog-go"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	service := grpc.NewService()
	service.Init()
	srv := NewService("pg.srv.entity.playerloginlog", service.Client())

	// Add
	resultAdd, err := srv.Add(context.TODO(), &AddRequest{
		PlayerID:   "1",
		OperatorID: "1",
		StartDate:  ptypes.TimestampNow(),
		EndDate:    nil,
	})

	if err != nil {
		fmt.Println("Add err:", err)
	} else {
		fmt.Println("Add result:", resultAdd.Message)
	}

	// Add 2
	// resultAdd, err = srv.Add(context.TODO(), &AddRequest{
	// 	PlayerID:   "1",
	// 	OperatorID: "1",
	// 	StartDate:  ptypes.TimestampNow(),
	// 	EndDate:    ptypes.TimestampNow(),
	// })

	// if err != nil {
	// 	fmt.Println("Add 2 err:", err)
	// } else {
	// 	fmt.Println("Add 2 result:", resultAdd.Message)
	// }

	// Get
	resultGet, err := srv.Get(context.TODO(), &GetRequest{PlayerID: "1", OperatorID: "1"})
	if err != nil {
		fmt.Println("Get err:", err)
	} else {
		fmt.Println("Get result:", resultGet)
	}
}
