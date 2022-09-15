package model

import "time"

type PlayerLoginLog struct {
	PlayerID   string     `bson:"playerID"`
	OperatorID string     `bson:"operatorID"`
	StartDate  time.Time  `bson:"startDate"`
	EndDate    *time.Time `bson:"endDate"`
}
