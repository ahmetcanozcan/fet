<h1 style="font-size:400%; letter-spacing: 12px" align="center">FET</h1>

<div style="margin-top: -20px" align="center">
  <strong>Build query dynamically</strong>
</div>

<div align="center">
  A <code> Go </code> library for building mongo queries with factory functions. 
</div>

<br/>

<div align="center">

![CI](https://github.com/ahmetcanozcan/fet/actions/workflows/ci.yaml/badge.svg)

</div>
<br>

<br />

## What is FET?

you can build queries with factory functions. It can make your repository functions more generic, and makes checking query variables easier with the built-in checker system.

## Installation

```bash
go get github.com/ahmetcanozcan/fet
```

## Getting Started

basic usage of the fet is building a query factory functions, and then this query can be used in mongo driver functions.

```golang
filter := fet.Build(
  fet.Field("username").Eq("test"),
)

err := collection.FindOne(ctx, filter)
```

instead of using `Field`, you can use `with functions` directly.

```golang
filter := fet.Build(
  fet.WithValueEq("userName", "test"),
)

// is equivalent of:
// filter := bson.M{
//   "userName": "test",
// }

err := collection.FindOne(ctx, filter)

```

## Features

### Generic Repository Functions

```golang
func (repo *repository) GetByFilters(ctx context.Context, filters ...[]fet.Updater) (*User, error) {
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

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
