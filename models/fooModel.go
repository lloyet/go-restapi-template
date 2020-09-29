package fooModel

import (
	"api/databases"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	Name string
	Age  int
	City string
}

func Search(filter interface{}) []*People {
	var results []*People
	cur, err := databases.FooCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil
	}
	for cur.Next(context.TODO()) {
		var elem People
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func FindById(id string) *People {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	var p People
	if err := databases.FooCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&p); err != nil {
		return nil
	}
	return &p
}

func Register(query People) interface{} {
	doc, err := databases.FooCollection.InsertOne(context.TODO(), query)
	if err != nil {
		return nil
	}
	return doc.InsertedID
}

func DeleteById(id string) bool {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	if _, err = databases.FooCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId}); err != nil {
		return false
	}
	return true
}
