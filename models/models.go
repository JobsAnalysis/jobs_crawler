package models

import (
	"context"
	"fmt"
	"jobs-crawler/pkg/setting"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	CreatedOn  int ` json:"createOn" bson:"createdOn"`
	ModifiedOn int ` json:"modifiedOn" bson:"modifiedOn"`
}

var client *mongo.Client

func Init() {
	sec, err := setting.Cfg.GetSection("database")
	dbName := sec.Key("NAME").String()
	host := sec.Key("HOST").String()
	dbType := sec.Key("TYPE").String()
	connectOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s/%s", dbType, host, dbName))
	client, err = mongo.Connect(context.Background(), connectOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func CloseDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
}
