package fet_test

import (
	"encoding/json"
	"fmt"

	"github.com/ahmetcanozcan/fet"
)

func ExampleBuild() {
	filter := fet.Build(
		fet.Field("name").Eq("dave"),
		fet.Field("age").Is(10),
	)

	b, _ := json.Marshal(filter)
	fmt.Println(string(b))
	// Output: {"age":10,"name":{"$eq":"dave"}}
}

func ExampleBuildSet() {
	query := fet.BuildSet(
		fet.Field("age").Is(18),
	)

	b, _ := json.Marshal(query)
	fmt.Println(string(b))
	// Output: {"$set":{"age":18}}
}

func ExampleChecker() {
	filter := fet.Build(
		fet.Field("username").Eq("dave", fet.IfNotEmpty),
		fet.Field("cardId").Eq("", fet.IfNotEmpty),
		fet.Field("car").Eq(nil, fet.IfNotNil),
	)

	b, _ := json.Marshal(filter)
	fmt.Println(string(b))
	// Output: {"username":{"$eq":"dave"}}

}

func ExampleStruct() {
	type User struct {
		Name  string `bson:"name"`
		Age   int
		Phone string
	}

	u := User{
		Name:  "Dave Bowman",
		Age:   18,
		Phone: "",
	}

	filter := fet.Struct(u).
		Pick("Name").
		Pick("Phone", fet.IfNotEmpty).
		Build()

	b, _ := json.Marshal(filter)
	fmt.Println(string(b))
	// Output: {"name":"Dave Bowman"}
}
