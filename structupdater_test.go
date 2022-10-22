package fet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name      string `bson:"name"`
	Age       int    `bson:"age"`
	Phone     string
	CreatedAt time.Time `bson:"created_at"`
	ZeroTime  time.Time
}

func Test_Struct_Pick(t *testing.T) {

	qt := assert.New(t)

	u := User{
		Name:      "John",
		Age:       30,
		Phone:     "",
		CreatedAt: time.Now(),
	}

	cases := []struct {
		updater  Updater
		expected M
	}{
		{
			updater: Struct(u).Pick("Name").Pick("Age").Pick("Phone"),
			expected: M{
				"name":  "John",
				"age":   30,
				"Phone": "",
			},
		},
		{
			updater: Struct(u).Pick("Name").Pick("NotField"),
			expected: M{
				"name": "John",
			},
		},
		{
			updater: Struct(u).Pick("CreatedAt").Pick("ZeroTime", IfNotZeroTime),
			expected: M{
				"created_at": u.CreatedAt,
			},
		},
		{
			updater: Struct(u).Pick("Name").Pick("Phone", IfNotEmpty),
			expected: M{
				"name": "John",
			},
		},
		{
			updater: Struct(u).PickAll(IfNotEmpty, IfNotZeroTime),
			expected: M{
				"name":       "John",
				"age":        30,
				"created_at": u.CreatedAt,
			},
		},
	}

	for _, c := range cases {
		m := M{}
		c.updater.Update(m)
		qt.Equal(c.expected, m)
	}
}

func Test_Struct_Build(t *testing.T) {
	qt := assert.New(t)

	s := struct {
		Name  string `bson:"name"`
		Age   int
		Phone string
	}{
		Name:  "Dave Bowman",
		Age:   18,
		Phone: "",
	}

	filter := Struct(s).
		Pick("Name").
		Pick("Phone", IfNotEmpty).
		Build()

	qt.Equal(s.Name, filter["name"])
	qt.Nil(filter["Phone"])
	qt.Nil(filter["Age"])

}

func Test_Struct(t *testing.T) {
	qt := assert.New(t)

	cases := []struct {
		value     interface{}
		errString string
	}{
		{User{}, ""},
		{&User{}, ""},
		{nil, ErrNilStruct.Error()},
		{"", ErrNotStruct.Error()},
		{1, ErrNotStruct.Error()},
		{1.1, ErrNotStruct.Error()},
	}

	for _, c := range cases {
		if c.errString == "" {
			Struct(c.value)
		} else {
			qt.PanicsWithError(c.errString, func() {
				Struct(c.value)
			})
		}
	}

}
