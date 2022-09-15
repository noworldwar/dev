package starter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PGITAb/bc-service-entity-playerloginlog-go/internal/model"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() {

	dataSource := fmt.Sprintf("mongodb://%s:%s",
		viper.GetString("db.host"),
		viper.GetString("db.port"))

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dataSource))

	if err != nil {
		log.Fatalln("Mongo err:", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("Mongo err:", err)
	}

	model.DBCollection = client.Database(viper.GetString("db.database")).Collection("PlayerLoginLog")
}
