// Package main implements an example usage of building update queries with fet
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
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	checkErr(err)

	collection := client.Database("test").Collection("test")

	_, err = collection.InsertOne(ctx, user{Name: "test"})
	checkErr(err)

	filter := fet.Build(
		fet.Field("name").Eq("test"),
	)

	query := fet.BuildSet(
		fet.Field("age").Is(30),
	)

	_, err = collection.UpdateOne(ctx, filter, query)
	checkErr(err)

	var u user
	err = collection.FindOne(ctx, filter).Decode(&u)
	checkErr(err)

	fmt.Println(u)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
