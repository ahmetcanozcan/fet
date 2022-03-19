package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmetcanozcan/fet"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URI = ""
)

type User struct {
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	CreatedAt time.Time `bson:"created_at"`
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	checkErr(err)

	collection := client.Database("test").Collection("test")

	u := User{
		Name:      "test",
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(ctx, u)
	checkErr(err)

	u.Age = 10

	filter := fet.Build(
		fet.Field("name").Eq("test"),
	)

	query := fet.BuildSet(
		fet.Struct(u).Pick("Age"),
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
