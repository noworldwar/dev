package starter

import (
	"log"
	"time"

	playerloginlog "github.com/PGITAb/bc-proto-entity-playerloginlog-go"
	"github.com/PGITAb/bc-service-entity-playerloginlog-go/internal/app/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

func GoMicroServerRun() {
	mService := grpc.NewService(
		micro.Name("pg.srv.entity.playerloginlog"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	playerloginlog.RegisterServiceHandler(mService.Server(), new(handler.ServiceHandler))

	if err := mService.Run(); err != nil {
		log.Fatal(err)
	}
}
