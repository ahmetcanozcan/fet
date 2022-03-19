package fet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Group(t *testing.T) {

	qt := assert.New(t)

	updaters := []Updater{
		Field("name").Is("Dave Bowman"),
		Field("age").Is(18),
	}

	group := Group(updaters...)

	bf := Build(updaters...)
	gf := Build(group)

	qt.Equal(bf, gf)
}
