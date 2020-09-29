package databases

import (
	"api/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

var FooCollection *mongo.Collection

func init() {
	/*Try Connection*/
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	/*Checking Connection*/
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	/*Storing Collections*/
	FooCollection = client.Database(config.MONGO_DB).Collection("peoples")
	fmt.Println("Connected to " + config.MONGO_PATH + "/" + config.MONGO_DB)
}
