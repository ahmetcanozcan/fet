// Package main implements an basic usage of fet
//
// this example shows that how you can build some filter
// without any other fet functionalities.
package main

import (
	"context"
	"fmt"

	"github.com/ahmetcanozcan/fet"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri = ""
)

type user struct {
	UserName string `bson:"userName"`
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	checkErr(err)

	collection := client.Database("test").Collection("test")

	_, err = collection.InsertOne(ctx, user{UserName: "test"})
	checkErr(err)

	var result user

	filter := fet.Build(
		fet.Field("userName").Eq("test"),
	)

	err = collection.FindOne(ctx, filter).Decode(&result)
	checkErr(err)
	fmt.Println(result)

	_, err = collection.DeleteOne(ctx, filter)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
