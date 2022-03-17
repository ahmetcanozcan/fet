package fet_test

import (
	"testing"

	"github.com/ahmetcanozcan/fet"
	"github.com/stretchr/testify/assert"
)

func TestGlobals(t *testing.T) {
	qt := assert.New(t)

	m := fet.M{}

	fet.Set(fet.Field("test_set").Eq("set")).Update(m)
	fet.SetOnInsert(fet.Field("test_insert").Eq("insert")).Update(m)
	fet.Unset(fet.Field("test_unset").Eq("unset")).Update(m)

	qt.Equal(m, fet.M{
		"$set": fet.M{
			"test_set": fet.M{
				"$eq": "set",
			},
		},
		"$setOnInsert": fet.M{
			"test_insert": fet.M{
				"$eq": "insert",
			},
		},
		"$unset": fet.M{
			"test_unset": fet.M{
				"$eq": "unset",
			},
		},
	})

}
