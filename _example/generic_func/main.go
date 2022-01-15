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

	user, err := getByFilter(collection, fet.Field("userName").Eq("test"))
	checkErr(err)
	fmt.Println(user)
}

func getByFilter(coll *mongo.Collection, filters ...fet.Updater) (user User, err error) {
	filter := fet.Build(filters...)
	err = coll.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
