package fet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_withFunc(t *testing.T) {

	qt := assert.New(t)

	m := M{}
	now := time.Now()

	withValueEq("key", "value").Update(m)
	withValueEq("key2", "", IfNotEmpty).Update(m)
	withValueEq("key3", 0, IfNotZero).Update(m)
	withValueEq("key4", time.Time{}, IfNotZeroTime).Update(m)
	withValueEq("key5", nil, IfNotNil).Update(m)
	withValueEq("key6", 24, IfNotNil).Update(m)
	withValueEq("key7", "", IfNotNil).Update(m)
	withValueEq("key8", now, IfNotNil).Update(m)
	withValueIs("key9", "test").Update(m)

	qt.Equal(m, M{
		"key": M{
			"$eq": "value",
		},
		"key6": M{
			"$eq": 24,
		},
		"key7": M{
			"$eq": "",
		},
		"key8": M{
			"$eq": now,
		},
		"key9": "test",
	})
}
