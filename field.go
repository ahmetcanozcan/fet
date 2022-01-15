package fet

type field struct {
	key string
}

func (f *field) Eq(value interface{}, checkers ...Checker) Updater {
	return WithValueEq(f.key, value, checkers...)
}

func (f *field) NotEq(value interface{}, checkers ...Checker) Updater {
	return WithValueNotEq(f.key, value, checkers...)
}

func (f *field) Exists(value interface{}, checkers ...Checker) Updater {
	return WithValueExists(f.key, value, checkers...)
}

func (f *field) In(value interface{}, checkers ...Checker) Updater {
	return WithValueIn(f.key, value, checkers...)
}

func (f *field) NotIn(value interface{}, checkers ...Checker) Updater {
	return WithValueNotIn(f.key, value, checkers...)
}

func (f *field) ElemMatch(value interface{}, checkers ...Checker) Updater {
	return WithValueElemMatch(f.key, value, checkers...)
}

func Field(key string) *field {
	return &field{key: key}
}
