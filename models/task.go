package models

import (
	"context"
	"fmt"
	"jobs-crawler/pkg/setting"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const taskCollectionName = "tasks"

type Task struct {
	Model
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name"  bson:"name"`
	Test       primitive.ObjectID `json:"test" bson:"test"`
	Position   string             `json:"position"  bson:"position"`
	CreatedBy  string             `json:"created_by" bson:"created_by"`
	ModifiedBy string             `json:"modified_by" bson:"modified_by"`
	Status     int                `json:"status" bson:"status"`
}

func GetTaks() (tasks []*Task) {
	coll := client.Database(setting.DbName).Collection(taskCollectionName)
	filter := bson.D{}
	var ctx = context.Background()
	docs, err := coll.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&docs)
	for docs.Next(ctx) {
		//Create a value into which the single document can be decoded
		var item Task
		// var elem Task
		err := docs.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(item)
		// bsonBytes, _ := bson.Marshal(item)
		// bson.Unmarshal(bsonBytes, &elem)
		tasks = append(tasks, &item)
	}
	return
}
