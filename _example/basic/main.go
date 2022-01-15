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
	UserName string `bson:"userName"`
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	checkErr(err)

	collection := client.Database("test").Collection("test")

	_, err = collection.InsertOne(ctx, User{UserName: "test"})
	checkErr(err)

	var result User

	filter := fet.Build(
		fet.Field("userName").Eq("test"),
	)

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
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
