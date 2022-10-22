// Package fet implements a field system that is built on withFunction,
//
// field system allows using MongoDB keywords and implementing some Checker logic easily.
// Under the hood  UpdateBuilders uses with functions
package fet

// UpdaterBuilder builds an updater with function calls same named as the keyword
type UpdaterBuilder interface {
	Is(value interface{}, checkers ...Checker) Updater
	Eq(value interface{}, checkers ...Checker) Updater
	NotEq(value interface{}, checkers ...Checker) Updater
	Exists(value interface{}, checkers ...Checker) Updater
	In(value interface{}, checkers ...Checker) Updater
	NotIn(value interface{}, checkers ...Checker) Updater
	ElemMatch(value interface{}, checkers ...Checker) Updater
	Gt(value interface{}, checkers ...Checker) Updater
	Gte(value interface{}, checkers ...Checker) Updater
	Lt(value interface{}, checkers ...Checker) Updater
	Lte(value interface{}, checkers ...Checker) Updater
}

type updaterBuilder struct {
	key string
}

func (f *updaterBuilder) Is(value interface{}, checkers ...Checker) Updater {
	return withValueIs(f.key, value, checkers...)
}

func (f *updaterBuilder) Eq(value interface{}, checkers ...Checker) Updater {
	return withValueEq(f.key, value, checkers...)
}

func (f *updaterBuilder) NotEq(value interface{}, checkers ...Checker) Updater {
	return withValueNotEq(f.key, value, checkers...)
}

func (f *updaterBuilder) Exists(value interface{}, checkers ...Checker) Updater {
	return withValueExists(f.key, value, checkers...)
}

func (f *updaterBuilder) In(value interface{}, checkers ...Checker) Updater {
	return withValueIn(f.key, value, checkers...)
}

func (f *updaterBuilder) NotIn(value interface{}, checkers ...Checker) Updater {
	return withValueNotIn(f.key, value, checkers...)
}

func (f *updaterBuilder) ElemMatch(value interface{}, checkers ...Checker) Updater {
	return withValueElemMatch(f.key, value, checkers...)
}

func (f *updaterBuilder) Gt(value interface{}, checkers ...Checker) Updater {
	return withValueGt(f.key, value, checkers...)
}

func (f *updaterBuilder) Gte(value interface{}, checkers ...Checker) Updater {
	return withValueGte(f.key, value, checkers...)
}

func (f *updaterBuilder) Lt(value interface{}, checkers ...Checker) Updater {
	return withValueLt(f.key, value, checkers...)
}

func (f *updaterBuilder) Lte(value interface{}, checkers ...Checker) Updater {
	return withValueLte(f.key, value, checkers...)
}

// Field returns UpdaterBuilder for given field name
func Field(key string) UpdaterBuilder {
	return &updaterBuilder{key: key}
}

// type check
var _ UpdaterBuilder = (*updaterBuilder)(nil)
