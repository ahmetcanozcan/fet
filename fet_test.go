package fet_test

import (
	"testing"

	"github.com/ahmetcanozcan/fet"
	"github.com/stretchr/testify/assert"
)

func Test_Build(t *testing.T) {

	qt := assert.New(t)

	filter := fet.Build(
		fet.Field("firstname").Eq("dave"),
		fet.Field("lastname").Eq("bowman"),
		fet.Or(
			fet.Field("age").Gt(30),
			fet.Field("age").Lt(10),
		),
	)

	qt.Equal(filter, fet.M{
		"firstname": fet.M{
			"$eq": "dave",
		},
		"lastname": fet.M{
			"$eq": "bowman",
		},
		"$or": []fet.M{
			{
				"age": fet.M{
					"$gt": 30,
				},
			},
			{
				"age": fet.M{
					"$lt": 10,
				},
			},
		},
	})

}

func Test_BuildSet(t *testing.T) {

	qt := assert.New(t)

	query := fet.BuildSet(
		fet.Field("firstname").Is("dave"),
		fet.Field("lastname").Is("bowman"),
	)

	qt.Equal(query, fet.M{
		"$set": fet.M{
			"firstname": "dave",
			"lastname":  "bowman",
		},
	})
}
