package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/PGITAb/bc-service-entity-playerloginlog-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Insert(m model.PlayerLoginLog) (string, error) {
	ctx, calcne := context.WithTimeout(context.Background(), 10*time.Second)
	defer calcne()

	cnt, _ := model.DBCollection.CountDocuments(ctx, bson.M{"playerID": m.PlayerID, "operatorID": m.OperatorID})
	if cnt > 0 {
		return "", fmt.Errorf("Data already exists")
	}

	res, err := model.DBCollection.InsertOne(ctx, m)
	return fmt.Sprintf("%v", res.InsertedID), err
}

func Get(playerID string, operatorID string) (*model.PlayerLoginLog, error) {
	m := new(model.PlayerLoginLog)
	err := model.DBCollection.FindOne(context.TODO(), bson.M{"playerID": playerID, "operatorID": operatorID}).Decode(&m)
	return m, err
}

func Update(m model.PlayerLoginLog) (int64, error) {
	ctx, calcne := context.WithTimeout(context.Background(), 10*time.Second)
	defer calcne()
	res, err := model.DBCollection.UpdateOne(ctx,
		bson.M{"playerID": m.PlayerID, "operatorID": m.OperatorID},
		bson.M{"$set": m})
	return res.ModifiedCount, err
}

func Delete(playerID string, operatorID string) (int64, error) {
	ctx, calcne := context.WithTimeout(context.Background(), 10*time.Second)
	defer calcne()
	res, err := model.DBCollection.DeleteOne(ctx, bson.M{"playerID": playerID, "operatorID": operatorID})
	return res.DeletedCount, err
}
