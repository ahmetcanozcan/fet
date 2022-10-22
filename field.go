// Package fet implements a field system that is built on withFunction,
//
// field system allows using MongoDB keywords and implementing some Checker logic easily
package fet

type field struct {
	key string
}

func (f *field) Is(value interface{}, checkers ...Checker) Updater {
	return withValueIs(f.key, value, checkers...)
}

func (f *field) Eq(value interface{}, checkers ...Checker) Updater {
	return withValueEq(f.key, value, checkers...)
}

func (f *field) NotEq(value interface{}, checkers ...Checker) Updater {
	return withValueNotEq(f.key, value, checkers...)
}

func (f *field) Exists(value interface{}, checkers ...Checker) Updater {
	return withValueExists(f.key, value, checkers...)
}

func (f *field) In(value interface{}, checkers ...Checker) Updater {
	return withValueIn(f.key, value, checkers...)
}

func (f *field) NotIn(value interface{}, checkers ...Checker) Updater {
	return withValueNotIn(f.key, value, checkers...)
}

func (f *field) ElemMatch(value interface{}, checkers ...Checker) Updater {
	return withValueElemMatch(f.key, value, checkers...)
}

func (f *field) Gt(value interface{}, checkers ...Checker) Updater {
	return withValueGt(f.key, value, checkers...)
}

func (f *field) Gte(value interface{}, checkers ...Checker) Updater {
	return withValueGte(f.key, value, checkers...)
}

func (f *field) Lt(value interface{}, checkers ...Checker) Updater {
	return withValueLt(f.key, value, checkers...)
}

func (f *field) Lte(value interface{}, checkers ...Checker) Updater {
	return withValueLte(f.key, value, checkers...)
}

func Field(key string) *field {
	return &field{key: key}
}
