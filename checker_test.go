package fet_test

import (
	"testing"
	"time"

	"github.com/ahmetcanozcan/fet"
	"github.com/stretchr/testify/assert"
)

func Test_IfNotNil(t *testing.T) {
	qt := assert.New(t)

	qt.False(fet.IfNotNil.Check("key", nil))
	qt.True(fet.IfNotNil.Check("key", ""))
	qt.True(fet.IfNotNil.Check("key", 0))
}

func Test_IfNotEmpty(t *testing.T) {
	qt := assert.New(t)

	qt.True(fet.IfNotEmpty.Check("key", nil))
	qt.False(fet.IfNotEmpty.Check("key", ""))
	qt.True(fet.IfNotEmpty.Check("key", 0))
}

func Test_IfNotZero(t *testing.T) {
	qt := assert.New(t)

	qt.True(fet.IfNotZero.Check("key", nil))
	qt.True(fet.IfNotZero.Check("key", ""))
	qt.False(fet.IfNotZero.Check("key", 0))
}

func Test_IfNotZeroTime(t *testing.T) {
	qt := assert.New(t)

	qt.True(fet.IfNotZeroTime.Check("key", 0)) // true because 0 is not time.Time
	qt.False(fet.IfNotZeroTime.Check("key", time.Time{}))
	qt.True(fet.IfNotZeroTime.Check("key", time.Now()))

}

func Test_CheckerFunc(t *testing.T) {
	qt := assert.New(t)

	checker := fet.CheckerFunc(func(key string, value interface{}) bool {
		return value != nil || key == "test"
	})

	qt.True(checker.Check("test", nil))
	qt.True(checker.Check("test", ""))
	qt.False(checker.Check("false", nil))
}
