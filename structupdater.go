package fet

import (
	"errors"
	"reflect"
)

var (
	ErrNilStruct    = errors.New("fet: nil pointer passed to Struct")
	ErrNotStruct    = errors.New("fet: non-struct passed to Struct")
	ErrInvalidField = errors.New("fet: invalid field")
)

type StructUpdater struct {
	v reflect.Value
	t reflect.Type
	m M
}

func (s *StructUpdater) Pick(name string, checkers ...Checker) *StructUpdater {
	vf := s.v.FieldByName(name)

	f, ok := s.t.FieldByName(name)

	if !ok {
		panic(ErrInvalidField)
	}

	val := vf.Interface()

	key := getKeyFromField(f, name)

	if !checkFilter(key, val, checkers...) {
		return s
	}

	if key == "" {
		key = name
	}

	s.m[key] = val

	return s
}

func (s *StructUpdater) PickAll(checkers ...Checker) *StructUpdater {
	for i := 0; i < s.t.NumField(); i++ {
		s.Pick(s.t.Field(i).Name, checkers...)
	}

	return s
}

func (s *StructUpdater) Update(m M) {
	for k, v := range s.m {
		m[k] = v
	}
}

func Struct(s interface{}) *StructUpdater {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if s == nil {
		panic(ErrNilStruct)
	}

	kind := v.Kind()
	if kind != reflect.Ptr && kind != reflect.Struct {
		panic(ErrNotStruct)
	}

	return &StructUpdater{
		v: v,
		t: t,
		m: M{},
	}
}
