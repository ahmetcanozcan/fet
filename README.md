<h1 style="font-size:400%; letter-spacing: 12px" align="center">FET</h1>

<div style="margin-top: -20px" align="center">
  <strong>Build queries dynamically</strong>
</div>

<div align="center">
well-tested <code> Go </code> library that allows you to build queries in dynamic, and programmatic way. 
</div>

<br/>

<div align="center">

![CI](https://github.com/ahmetcanozcan/fet/actions/workflows/ci.yaml/badge.svg)

</div>
<br>

<br />

## What is FET?

Fet is a library that allows you to build queries in dynamic, and programmatic way. You can build mongo queries with function calls, instead of writing nested maps. And you have a unique, built-in logical control on your queries easily.

## Installation

```bash
go get github.com/ahmetcanozcan/fet
```

## Getting Started

Main functionality of FET is to build queries with function calls using `fet.Build`. This function takes a list of `fet.Updater` and apply this updaters into a new query that can be used on mongo driver functions.

```golang
filter := fet.Build(
  fet.Field("name").Eq("dave"),
)

err := collection.FindOne(ctx, filter)


```

You can build set queries as well, using `fet.BuildSet`.

```golang

filter := fet.Build(
  fet.Field("name").Eq("dave"),
)

query := fet.BuildSet(
  fet.Field("age").Is(18),
)

_, err := collection.UpdateOne(ctx, filter, query)

// is equivalent to

filter := bson.M{
  "name": bson.M{
    "$eq": "dave",
  },
}

query := bson.M{
  "$set": bson.M{
    "age": 18,
  },
}

err := collection.FindOne(ctx, filter, query)

```

## Features

### Generic Repository Functions

```golang
func (repo *repository) GetByFilters(ctx context.Context, filters ...fet.Updater) (*User, error) {
  // ...
  filter := fet.Build(filters...)
  err := collection.FindOne(ctx, filter).Decode(&user)
  // ...
  return user, err
}

// This function can be called
user, err := repo.GetByFilters(
  ctx,
  fet.WithValueEq("username", "test"),
  fet.WithValueEq("password", "test"),
)

user, err := repo.GetByFilters(
  ctx,
  fet.Field("username").Eq("test"),
  fet.Field("password").Eq("test"),
  fet.Or(
    fet.Field("pretty").Eq(true),
    fet.Field("admin").Eq(true),
  )
)

```

### Checker

checker is basic function that returns true depends on field name and query variable. `FET` has some built-in checker functions such as `IfNotNil`, `IfNotEmpty`, `IfNotZeroTime`.

```golang
filter := fet.Build(
  fet.Field("username").Eq(username),
  fet.Field("cardId").Eq(cardId, fet.IfNotEmpty),
)

// filter is change depends on cardId

// if cardId is empty, filter is equivalent of:
filter := bson.M{
  "username": username,
}

// if cardId is not empty, filter is equivalent of:
filter := bson.M{
  "username": username,
  "cardId": cardId,
}

```

### Struct Picking

if your code reads fields from a struct, you can use the `Struct` functionality to shortened your code.

```golang

type User struct {
  Name string `bson:"name"`
  Age int
  Phone string
}

u := &User{
  Name: "Dave Bowman",
  Age: 18,
  Phone: "",
}

f := fet.Build(
  fet.Field("name").Is(u.Name),
  fet.Field("Phone").Is(u.Phone, fet.IfNotEmpty)
)

f := fet.Struct(u).
  Pick("Name").
  Pick("Phone", fet.IfNotEmpty).
  Build()

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
