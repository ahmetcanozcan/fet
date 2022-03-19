package fet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logicalOperatioWithKeyword(t *testing.T) {
	qt := assert.New(t)

	m := M{}

	const keyword = "keyword"

	eqUpdater := withValueEq("key", "value")
	existsUpdater := withValueEq("key", "value")

	logicalOperatioWithKeyword(
		keyword,
		eqUpdater,
		existsUpdater,
	).Update(m)

	eqM := M{}
	eqUpdater.Update(eqM)
	existsM := M{}
	existsUpdater.Update(existsM)

	qt.Equal(M{
		keyword: []M{
			eqM,
			existsM,
		},
	}, m)

}

func Test_Or(t *testing.T) {
	qt := assert.New(t)

	m := M{}

	eqUpdater := withValueEq("key", "value")
	existsUpdater := withValueEq("key", "value")

	Or(
		eqUpdater,
		existsUpdater,
	).Update(m)

	eqM := M{}
	eqUpdater.Update(eqM)
	existsM := M{}
	existsUpdater.Update(existsM)

	qt.Equal(M{
		KeywordOr: []M{
			eqM,
			existsM,
		},
	}, m)

}
