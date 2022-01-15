package fet_test

import (
	"testing"
	"time"

	"github.com/ahmetcanozcan/fet"
	"github.com/stretchr/testify/assert"
)

func Test_WithFuncEq(t *testing.T) {

	qt := assert.New(t)

	m := fet.M{}
	now := time.Now()

	fet.WithValueEq("key", "value").Update(m)
	fet.WithValueEq("key2", "", fet.IfNotEmpty).Update(m)
	fet.WithValueEq("key3", 0, fet.IfNotZero).Update(m)
	fet.WithValueEq("key4", time.Time{}, fet.IfNotZeroTime).Update(m)
	fet.WithValueEq("key5", nil, fet.IfNotNil).Update(m)
	fet.WithValueEq("key6", 24, fet.IfNotNil).Update(m)
	fet.WithValueEq("key7", "", fet.IfNotNil).Update(m)
	fet.WithValueEq("key8", now, fet.IfNotNil).Update(m)

	qt.Equal(m, fet.M{
		"key": fet.M{
			"$eq": "value",
		},
		"key6": fet.M{
			"$eq": 24,
		},
		"key7": fet.M{
			"$eq": "",
		},
		"key8": fet.M{
			"$eq": now,
		},
	})
}
