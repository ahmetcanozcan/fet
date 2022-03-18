package main

import (
	"context"
	"fmt"

	"github.com/ahmetcanozcan/fet"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URI = ""
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	checkErr(err)

	collection := client.Database("test").Collection("test")

	_, err = collection.InsertOne(ctx, User{Name: "test"})
	checkErr(err)

	filter := fet.Build(
		fet.Field("name").Eq("test"),
	)

	query := fet.BuildSet(
		fet.Field("age").Is(30),
	)

	_, err = collection.UpdateOne(ctx, filter, query)
	checkErr(err)

	var user User
	err = collection.FindOne(ctx, filter).Decode(&user)
	checkErr(err)

	fmt.Println(user)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
