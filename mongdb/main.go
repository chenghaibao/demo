package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// https://docs.mongodb.com/drivers/go/current/fundamentals/crud/read-operations/retrieve/
// https://studygolang.com/articles/16846
// https://mongoing.com/archives/27257
// https://pkg.go.dev/github.com/mongodb/mongo-go-driver#section-readme
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"),
		options.Client().SetAuth(options.Credential{Username: "admin", Password: "admin"}))
	collection := client.Database("test").Collection("hb")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//res, _ := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3}})

	//var result struct {
	//	Value float64
	//}
	//filter := bson.M{"name": "pi"}
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	//err := collection.FindOne(ctx, filter).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		//var result bson.M
		var result struct {
			Id    primitive.ObjectID `bson:"_id"`
			Name  string             `bson:"name"`
			Value float64            `bson:"value"`
		}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println(result.Id.Hex())
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
