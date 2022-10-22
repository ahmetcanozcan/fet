// Package main implements an example usage of generic functions with fet
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

	u, err := getByFilter(collection, fet.Field("userName").Eq("test"))
	checkErr(err)
	fmt.Println(u)
}

func getByFilter(coll *mongo.Collection, filters ...fet.Updater) (u user, err error) {
	filter := fet.Build(filters...)
	err = coll.FindOne(context.TODO(), filter).Decode(&u)
	return u, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
