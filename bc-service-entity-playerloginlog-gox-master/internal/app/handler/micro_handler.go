package handler

import (
	"context"
	"log"

	playerloginlog "github.com/PGITAb/bc-proto-entity-playerloginlog-go"
	"github.com/PGITAb/bc-service-entity-playerloginlog-go/internal/model"
	"github.com/golang/protobuf/ptypes"
)

type ServiceHandler struct {
}

func (h *ServiceHandler) Add(ctx context.Context, in *playerloginlog.AddRequest, out *playerloginlog.AddResponse) error {
	sdate, _ := ptypes.Timestamp(in.StartDate)
	m := model.PlayerLoginLog{
		PlayerID:   in.PlayerID,
		OperatorID: in.OperatorID,
		StartDate:  sdate,
	}

	if in.EndDate != nil {
		edate, _ := ptypes.Timestamp(in.EndDate)
		m.EndDate = &edate
	}

	_, err := Insert(m)

	if err != nil {
		log.Println("Insert err:", err)
		out.Message = err.Error()
	} else {
		log.Println("Insert ok:", in)
		out.Message = "ok"
	}

	// affected, err := Update(m)

	// if err != nil {
	// 	log.Println("Update err:", err)
	// 	out.Message = err.Error()
	// } else if affected == 0 {
	// 	log.Println("Update Data not found or unchanged:", in)
	// 	out.Message = "Data not found or unchanged"
	// } else {
	// 	log.Println("Update ok:", in)
	// 	out.Message = "ok"
	// }

	return nil
}

func (h *ServiceHandler) Get(ctx context.Context, in *playerloginlog.GetRequest, out *playerloginlog.GetResponse) error {
	data, err := Get(in.PlayerID, in.OperatorID)
	if err != nil {
		log.Println("Get err:", err)
		return err
	}

	// log.Println("Get:", data)

	out.PlayerID = data.PlayerID
	out.OperatorID = data.OperatorID
	out.StartDate, _ = ptypes.TimestampProto(data.StartDate)
	if data.EndDate != nil {
		out.EndDate, _ = ptypes.TimestampProto(*data.EndDate)
	}

	// affected, err := Delete(in.PlayerID, in.OperatorID)
	// if err != nil {
	// 	log.Println("Delete err:", err)
	// } else if affected == 0 {
	// 	log.Println("Delete Data not found:", in)
	// } else {
	// 	log.Println("Delete ok:", in)
	// }

	return nil
}
